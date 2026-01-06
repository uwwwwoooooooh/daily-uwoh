package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/uwwwwoooooooh/daily-uwoh/internal/db/sqlc"
)

// Store provides all functions to execute db queries and transactions
type Store interface {
	ArtworkRepository
	UserRepository
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	conn *pgxpool.Pool
	sqlc.Store
}

// NewStore creates a new store
func NewStore(conn *pgxpool.Pool) Store {
	return &SQLStore{
		conn:  conn,
		Store: sqlc.NewStore(conn),
	}
}
