package main

import (
	"flag"
	"log/slog"
	"os"
)

// config stores all the configuration needed for running web application
type config struct {
	port int    // specifies the port no on which application listens
	env  string // specifies the mode (development|production) application runs in
}

// application stores all the dependencies handlers, middleware etc  in application needs
type application struct {
	config config
	logger *slog.Logger // structured logger for all logging needs of application
}

func main() {
	cfg := config{}

	flag.IntVar(&cfg.port, "port", 4000, "Web server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|production)")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	err := app.serve()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
