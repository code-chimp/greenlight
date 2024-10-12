package main

import (
	"flag"
	"fmt"
	"github.com/code-chimp/greenlight/internal/vcs"
	"log/slog"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

var revision = vcs.Revision()

// config struct holds the configuration settings for the application.
type config struct {
	port int
	env  string
}

// application holds the application-wide dependencies.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	displayVersion := flag.Bool("version", false, "Display version information")
	flag.Parse()

	if *displayVersion {
		fmt.Printf("Greenlight Website:\n\tVersion:\t%s\n\tRevision:\t%s\n", version, revision)
		os.Exit(0)
	}

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", fmt.Sprintf("http://localhost%s", srv.Addr), "env", cfg.env)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
