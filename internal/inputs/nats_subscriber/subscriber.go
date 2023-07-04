package nats_subscriber

import (
	"fmt"

	"l0_wb/internal/app"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

type Settings struct {
	ClusterID   string
	ClientID    string
	Subject     string
	DurableName string
	NatsHost    string
	NatsPort    string
}

type subscriber struct {
	subj        string
	durableName string
	conn        stan.Conn
	subs        stan.Subscription
	app         app.Service
}

func New(settings Settings, app app.Service) (subscriber, error) {
	connString := fmt.Sprintf("nats://%s:%s", settings.NatsHost, settings.NatsPort)
	conn, err := stan.Connect(
		settings.ClusterID,
		settings.ClientID,
		stan.NatsURL(connString),
		stan.NatsOptions(nats.MaxReconnects(1<<63-1)),
	)
	if err != nil {
		return subscriber{}, fmt.Errorf("can't create connection to %s: %v", connString, err)
	}

	return subscriber{
		subj:        settings.Subject,
		durableName: settings.DurableName,
		app:         app,
		conn:        conn,
	}, nil
}
