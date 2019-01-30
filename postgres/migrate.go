package postgres

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

	for _, migration := range migrations {
		_, err := conn.Exec(migration.Up)
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
