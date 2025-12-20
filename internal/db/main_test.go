package db

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc"
)

var testQueries *sqlc.Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	// TODO: Replace with your actual database DSN or load from env
	dsn := "postgres://user:password@localhost:5432/dbname?sslmode=disable"
	if envDSN := os.Getenv("DATABASE_URL"); envDSN != "" {
		dsn = envDSN
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	testDB, err = pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}
	defer testDB.Close()

	if err := testDB.Ping(ctx); err != nil {
		log.Fatalf("cannot ping db: %v", err)
	}

	testQueries = sqlc.New(testDB)

	os.Exit(m.Run())
}
