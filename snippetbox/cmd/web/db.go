package main

import (
  "context"
	"github.com/jackc/pgx/v5/pgxpool"
)

func createConnectionPool(dsn string) (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	if err := dbpool.Ping(context.Background()); err != nil {
		return nil, err
	}
	return dbpool, nil
}
