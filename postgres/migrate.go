package postgres

import (
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

// Migrate .
func Migrate() error {
	wrapErr := func(err error) error {
		return errors.Wrap(err, "Migrating")
	}

	err := EnsureDatabaseExists("gonference")
	if err != nil {
		return wrapErr(err)
	}

	migrations, err := GetMigrations()
	if err != nil {
		return wrapErr(err)
	}

	conn, err := GetConnectionForDatabase("gonference")
	if err != nil {
		return wrapErr(err)
	}

	err = EnsureMigrationHistoryTableExists(conn)
	if err != nil {
		return wrapErr(err)
	}

	lastMigrationID, err := GetMostRecentMigrationID(conn)
	if err != nil {
		return wrapErr(err)
	}

	for _, migration := range migrations {
		if migration.Order <= lastMigrationID {
			continue
		}
		log.Println("Applying migration [" + strconv.Itoa(migration.Order) + "] '" + migration.Name + "'")
		_, err := conn.Exec(migration.Up)
		if err != nil {
			return wrapErr(err)
		}
		err = RegisterMigration(conn, migration)
		if err != nil {
			return wrapErr(err)
		}
	}

	return nil
}

// EnsureDatabaseExists .
func EnsureDatabaseExists(name string) error {
	wrapErr := func(err error) error {
		return errors.Wrap(err, "Ensuring database '"+name+"' exists")
	}

	conn, err := GetConnection()
	if err != nil {
		return wrapErr(err)
	}

	exists, err := CheckDatabaseExists(conn, name)
	if err != nil {
		return wrapErr(err)
	}

	if !exists {
		log.Println("Database 'gonference' doesn't exist. Creating.")
		_, err = conn.Exec("CREATE DATABASE gonference;")
		if err != nil {
			return wrapErr(err)
		}
	} else {
		log.Println("Database 'gonference' already exists.")
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

// GetMostRecentMigrationID .
func GetMostRecentMigrationID(db *sqlx.DB) (int, error) {
	wrapErr := func(err error) error {
		return errors.Wrap(err, "Getting most recent migration ID")
	}

	migrations := []int{}
	err := db.Select(&migrations, "SELECT id FROM __migration_history order by id desc LIMIT 1")
	if err != nil {
		return -1, wrapErr(err)
	}
	if len(migrations) == 0 {
		return -1, nil
	}
	return migrations[0], nil
}

// CheckDatabaseExists .
func CheckDatabaseExists(db *sqlx.DB, name string) (bool, error) {
	wrapErr := func(err error) error {
		return errors.Wrap(err, "Checking database '"+name+"' exists")
	}

	rows, err := db.Query("SELECT 1 FROM pg_database WHERE datname=$1", name)
	if err != nil {
		return false, wrapErr(err)
	}

	count := 0
	for rows.Next() {
		count++
	}

	return count == 1, nil
}
