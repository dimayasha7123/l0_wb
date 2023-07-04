package http_server

import (
	"net/http"
	"time"

	"l0_wb/internal/app"
	"l0_wb/internal/inputs/http_server/handlers"

	"github.com/gorilla/mux"
)

func New(app app.Service, addr string) *http.Server {
	router := mux.NewRouter()
	router.Use(loggingMiddleware)

	subrouter := router.PathPrefix("/api/v1").Subrouter()
	// api/v1/model/<id>
	subrouter.Handle("/model/{id:[A-Za-z0-9]+}", handlers.NewModelHandler(app)).Methods(http.MethodGet)

	router.NotFoundHandler = router.NewRoute().HandlerFunc(http.NotFound).GetHandler()

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
		IdleTimeout:  time.Second * 10,
	}

	return srv
}
