package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Include PostgreSQL library
)

// GetConnection .
func GetConnection(connectionString string) (*sqlx.DB, error) {
	return sqlx.Connect("postgres", connectionString)
}
