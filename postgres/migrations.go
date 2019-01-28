package postgres

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

// Migration .
type Migration struct {
	Order int
	Name  string
	Up    string
	Down  string
}

// MigrationFile .
type MigrationFile struct {
	Order     int
	Name      string
	Direction string
	Path      string
}

// GetMigrations .
func GetMigrations() ([]Migration, error) {
	var files []MigrationFile
	migrations := make([]Migration, 0)
	migrationsIndex := make(map[string]*Migration)

	if f, err := getMigrationsFiles(); err == nil {
		files = f
	} else {
		return nil, err
	}

	for _, file := range files {
		var migration *Migration
		if value, ok := migrationsIndex[file.Name]; ok {
			migration = value
		} else {
			migration = &Migration{Name: file.Name, Order: file.Order}
			migrationsIndex[file.Name] = migration
		}

		if file.Direction == "up" {
			migration.Up = file.Path
		} else {
			migration.Down = file.Path
		}
	}

	for key := range migrationsIndex {
		migrations = append(migrations, *migrationsIndex[key])
	}

	sort.Slice(migrations, func(a, b int) bool {
		return migrations[a].Order < migrations[b].Order
	})

	return migrations, nil
}

func getMigrationsFiles() ([]MigrationFile, error) {
	migrationsDirectoryPath, err := filepath.Abs("./postgres/migrations")
	if err != nil {
		return nil, err
	}

	migrationFiles := make([]MigrationFile, 0)
	files, err := ioutil.ReadDir(migrationsDirectoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		mf := parseMigrationFileName(file.Name())
		mf.Path = filepath.Join(migrationsDirectoryPath, file.Name())
		migrationFiles = append(migrationFiles, mf)
	}

	return migrationFiles, nil
}

func parseMigrationFileName(name string) MigrationFile {
	parts := strings.Split(name, ".")
	order, err := strconv.Atoi(parts[0])
	if err != nil {
		fmt.Println(err)
	}
	return MigrationFile{
		Order:     order,
		Name:      parts[1],
		Direction: parts[2],
	}
}
