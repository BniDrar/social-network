package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3" // Add this import
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"

	"socialNetwork/pkg/config"
	"socialNetwork/pkg/loger"
)

// init database
func InitDB(dbConf config.Database) (*sql.DB, error) {
	db, err := sql.Open(dbConf.Driver, dbConf.FileName)
	if err != nil {
		return nil, fmt.Errorf("could not open database: %w", err)
	}
	//  run migrations
	if err := RunMigrations(); err != nil {
		return nil, fmt.Errorf("could not run migrations: %w", err)
	}
	loger.NewLogger().Info.Println("Database connected")
	return db, nil
}

// MigrateDB migrates the database to the latest version
func RunMigrations() error {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	workDir, err := os.Getwd()
	loger.NewLogger().Info.Println("Working directory: ", workDir)
	if err != nil {
		return fmt.Errorf("failed to get working directory: %w", err)
	}
	sourceURL := fmt.Sprintf("file://%s/%s", workDir, config.Database.SchemeDir)
	// Open the database connection
	db, err := sql.Open(config.Database.Driver, config.Database.FileName)
	if err != nil {
		return fmt.Errorf("could not open database: %w", err)
	}
	defer db.Close()

	// Create the SQLite driver instance
	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		return fmt.Errorf("could not create driver instance: %w", err)
	}

	// Create a new migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		sourceURL,
		config.Database.Driver,
		driver,
	)
	if err != nil {
		return fmt.Errorf("could not create migrate instance: %w", err)
	}
	err = m.Force(11) // 11 is migrations version number, you may use your latest version
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("could not apply migrations: %w", err)
	}
	return nil
}
