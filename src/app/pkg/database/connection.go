package database

import (
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
