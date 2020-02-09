package migrator

import (
	"gonference/pkg/infrastructure/database/client"
	"gonference/pkg/infrastructure/logger"
	"gonference/pkg/utils"
)

func Migrate(conn *client.DBClient) {
	log := logger.Instance

	migrations := getMigrations()
	ensureMigrationHistoryTableExists(conn)

	lastMigrationID := getMostRecentMigrationID(conn)

	for _, migration := range migrations {
		if migration.Order <= lastMigrationID {
			log.Info("Migration", migration, "already applied. Skipping.")
			continue
		}
		log.Info("Applying migration ", migration)
		conn.Exec(migration.Up)
		registerMigration(conn, migration)
	}

	log.Info("Migration done.")
}

func ensureMigrationHistoryTableExists(conn *client.DBClient) {
	conn.Exec(`
		CREATE TABLE IF NOT EXISTS __migration_history (
			id INTEGER PRIMARY KEY,
			name VARCHAR (200) NOT NULL
		);
	`)
}

func registerMigration(conn *client.DBClient, migration Migration) {
	conn.Exec(`
		INSERT INTO __migration_history
		("id", "name") VALUES ($1, $2);
	`, migration.Order, migration.Name)
}

func getMostRecentMigrationID(conn *client.DBClient) int {
	rows := conn.Query("SELECT id FROM __migration_history order by id desc LIMIT 1")

	migrations := make([]int, 0)
	for rows.Next() {
		id := 0
		err := rows.Scan(&id); utils.Check(err)
		migrations = append(migrations, id)
	}

	if len(migrations) == 0 {
		return -1
	}

	return migrations[0]
}
