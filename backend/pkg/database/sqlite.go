package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3" // Add this import
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"socialNetwork/pkg/config"
)

// init database
func InitDB(dbConf config.Database) (*sql.DB, error) {
	db, err := sql.Open(dbConf.Driver, dbConf.FileName)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}
	//  run migrations
	if err := RunMigrations(dbConf, db); err != nil {
		return nil, fmt.Errorf("could not run migrations: %w", err)
	}
	return db, nil
}

// MigrateDB migrates the database to the latest version
func RunMigrations(Database config.Database, db *sql.DB) error {
	workDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	sourceURL := fmt.Sprintf("file://%s/%s", workDir, Database.SchemeDir)
	// Open the database connection
	if err != nil {
		return fmt.Errorf("could not open database: %w", err)
	}

	// Create the SQLite driver instance
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not create driver instance: %w", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		Database.Driver,
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not apply migrations: %w", err)
	}
	return nil
}
