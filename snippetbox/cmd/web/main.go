package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable", "PostgreSQL data source name")
	flag.Parse()

	loggerOpts := &slog.HandlerOptions{
		AddSource: true,
	}
	logger := slog.New(slog.NewJSONHandler(os.Stderr, loggerOpts))

	dbpool, err := createConnectionPool(*dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to ping database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	app := &Application{
		logger: logger,
	}

	mux := app.routes()

	logger.Info("starting server", "addr", *addr)

	err = http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
