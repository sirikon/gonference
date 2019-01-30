package postgres

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Include PostgreSQL library
)

// GetConnection .
func GetConnection() (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "user=postgres password=12345 sslmode=disable")
}

// GetConnectionForDatabase .
func GetConnectionForDatabase(databaseName string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", "user=postgres password=12345 dbname="+databaseName+" sslmode=disable")
}
