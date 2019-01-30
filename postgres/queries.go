package postgres

import (
	"github.com/jmoiron/sqlx"
)

// CheckDatabaseExists .
func CheckDatabaseExists(db *sqlx.DB, name string) (bool, error) {
	rows, err := db.Query("SELECT 1 FROM pg_database WHERE datname=$1", name)
	if err != nil {
		return false, err
	}

	count := 0
	for rows.Next() {
		count++
	}

	return count == 1, nil
}
