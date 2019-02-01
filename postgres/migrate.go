package postgres

import (
	"github.com/jmoiron/sqlx"
)

// Migrate .
func Migrate() error {
	err := EnsureDatabaseExists("gonference")
	if err != nil {
		return err
	}

	migrations, err := GetMigrations()
	if err != nil {
		return err
	}

	conn, err := GetConnectionForDatabase("gonference")
	if err != nil {
		return err
	}

	err = EnsureMigrationHistoryTableExists(conn)
	if err != nil {
		return err
	}

	for _, migration := range migrations {
		_, err := conn.Exec(migration.Up)
		if err != nil {
			return err
		}
		err = RegisterMigration(conn, migration)
		if err != nil {
			return err
		}
	}

	return nil
}

// EnsureDatabaseExists .
func EnsureDatabaseExists(name string) error {
	conn, err := GetConnection()
	if err != nil {
		return err
	}

	exists, err := CheckDatabaseExists(conn, name)
	if err != nil {
		return err
	}

	if !exists {
		_, err = conn.Exec("CREATE DATABASE gonference;")
		if err != nil {
			return err
		}
	}

	return nil
}

// EnsureMigrationHistoryTableExists .
func EnsureMigrationHistoryTableExists(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS __migration_history (
			id INTEGER PRIMARY KEY,
			name VARCHAR (200) NOT NULL
		);
	`)
	return err
}

// RegisterMigration .
func RegisterMigration(db *sqlx.DB, migration Migration) error {
	_, err := db.Exec(`
		INSERT INTO __migration_history
		("id", "name") VALUES ($1, $2);
	`, migration.Order, migration.Name)
	return err
}

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
