package database

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Include PostgreSQL library
)

// GetConnection .
func GetConnection(connectionString string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", connectionString)
	if err != nil { return nil, err }
	db.SetMaxOpenConns(90)
	return db, nil
}

func GetConnectionPool(connectionString string) (*pgxpool.Pool, error) {
	pool, err := pgxpool.Connect(context.Background(), connectionString)
	if err != nil { return nil, err }
	return pool, nil
}
