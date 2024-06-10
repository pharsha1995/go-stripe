package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// server defines the properties of http.Server and starts it.
// returns an error if any problem is faced by server
func (app *application) serve() error {
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(app.logger.Handler(), slog.LevelError),
	}

	app.logger.Info("starting server", "addr", server.Addr, "mode", app.config.env)

	return server.ListenAndServe()
}
