package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"l0_wb/internal/adapters/cache"
	"l0_wb/internal/adapters/json_validator"
	"l0_wb/internal/adapters/repository"
	"l0_wb/internal/app"
	"l0_wb/internal/inputs/http_server"
	"l0_wb/internal/inputs/nats_subscriber"
	"l0_wb/internal/utils/logger"

	"go.uber.org/zap"
)

func main() {
	var (
		modelSchemaPath = "model/schema.json"
		postgresDSN     = "host=db port=5432 user=l0_back_user password=l0_back_hard_password dbname=l0_back_db sslmode=disable"
		httpHost        = "0.0.0.0"
		httpPort        = "8000"
		natsClusterID   = "cluster"
		natsClientID    = "client_0"
		natsSubject     = "models"
		natsDurableName = "d_models"
		natsHost        = "nats_streaming"
		natsPort        = "4223"
	)

	flag.StringVar(&modelSchemaPath, "schema", modelSchemaPath, "path to model schema file")
	flag.StringVar(&postgresDSN, "postgres_dsn", postgresDSN, "postgres connection string")
	flag.StringVar(&httpHost, "host", httpHost, "http server host")
	flag.StringVar(&httpPort, "port", httpPort, "http server port")
	flag.StringVar(&natsClusterID, "cluster", natsClusterID, "nats cluster ID")
	flag.StringVar(&natsClientID, "client", natsClientID, "nats client ID")
	flag.StringVar(&natsSubject, "subj", natsSubject, "nats subject")
	flag.StringVar(&natsDurableName, "d_name", natsDurableName, "nats durable name")
	flag.StringVar(&natsHost, "nats_host", natsHost, "nats host")
	flag.StringVar(&natsPort, "nats_port", natsPort, "nats port")

	flag.Parse()

	zapLogger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't create logger: %v", err)
	}
	sugarLogger := zapLogger.Sugar()
	logger.SetLogger(sugarLogger)

	ctx := context.Background()

	schemaBytes, err := os.ReadFile(modelSchemaPath)
	if err != nil {
		logger.Log().Fatalf("can't read model schema file: %v", err)
	}

	validator, err := json_validator.New(string(schemaBytes))
	if err != nil {
		logger.Log().Fatalf("can't create json model validator: %v", err)
	}

	repo, err := repository.New(ctx, postgresDSN)
	if err != nil {
		logger.Log().Fatalf("can't create repository: %v", err)
	}

	service, err := app.New(ctx, cache.New(), repo, validator)
	if err != nil {
		logger.Log().Fatalf("can't create app service: %v", err)
	}
	logger.Log().Info("Create app service")

	settings := nats_subscriber.Settings{
		ClusterID:   natsClusterID,
		ClientID:    natsClientID,
		Subject:     natsSubject,
		DurableName: natsDurableName,
		NatsHost:    natsHost,
		NatsPort:    natsPort,
	}
	subs, err := nats_subscriber.New(settings, service)
	if err != nil {
		logger.Log().Fatalf("Can't create nats streaming subscriber: %v", err)
	}
	logger.Log().Info("Create nats streaming subscriber")

	err = subs.BeginListen(ctx)
	if err != nil {
		logger.Log().Fatalf("Can't begin listen nats stream: %v", err)
	}
	logger.Log().Info("Begin listen nats stream")

	httpServer := http_server.New(service, fmt.Sprintf("%s:%s", httpHost, httpPort))
	go func() {
		err := httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				logger.Log().Info("Server was closed")
			} else {
				logger.Log().Fatalf("Get error while serving: %v", err)
			}
		}
	}()
	logger.Log().Info("HTTP server listening")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	err = httpServer.Shutdown(ctx)
	if err != nil {
		logger.Log().Errorf("Can't shutdown http server")
	}
	subs.Close()
	logger.Log().Info("Subscription was closed")

	logger.Log().Info("All services were shutdown")
}
