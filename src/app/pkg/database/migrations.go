package database

import (
	"fmt"
	"gonference/pkg/assets"
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

func (m Migration) String() string {
	return "[" + strconv.Itoa(m.Order) + "] '" + m.Name + "'"
}

// MigrationFile .
type MigrationFile struct {
	Order     int
	Name      string
	Direction string
	Content   string
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
			migration.Up = file.Content
		} else {
			migration.Down = file.Content
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
	box := assets.DatabaseMigrations

	migrationFiles := make([]MigrationFile, 0)
	files := box.List()

	for _, file := range files {
		data, err := box.FindString(file)
		if err != nil {
			return nil, err
		}
		mf := parseMigrationFileName(file)
		mf.Content = data
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
