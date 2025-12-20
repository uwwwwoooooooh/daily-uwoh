package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

// ConnectDB initializes the database connection
func ConnectDB(dsn string) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	config.MaxConns = 100
	config.MinConns = 10

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(context.Background()); err != nil {
		return nil, err
	}

	log.Println("✅ Connected to Database")
	return pool, nil
}
