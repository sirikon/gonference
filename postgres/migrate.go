package postgres

// Migrate .
func Migrate() error {
	migrations, err := GetMigrations()
	if err != nil {
		return err
	}

	conn, err := GetConnection()
	if err != nil {
		return err
	}

	_, err = conn.Exec("CREATE DATABASE gonference;")
	if err != nil {
		return err
	}

	conn, err = GetConnectionForDatabase("gonference")

	for _, migration := range migrations {
		_, err := conn.Exec(migration.Up)
		if err != nil {
			return err
		}
	}

	return nil
}
