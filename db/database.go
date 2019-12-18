package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jboulet/fizzbuzz-go/config"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	_ "github.com/lib/pq"
)

func CreateDatabase() (*sql.DB, error) {

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.S.BD_HOST, config.S.DB_PORT, config.S.DB_USERNAME, config.S.DB_PASSWORD, config.S.DB_NAME)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	if err := migrateDatabase(db); err != nil {
		return db, err
	}

	return db, nil
}

func migrateDatabase(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	migration, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s/db/migrations", dir),
		"postgres",
		driver,
	)
	if err != nil {
		return err
	}

	migration.Log = &MigrationLogger{}

	migration.Log.Printf("Applying database migrations")
	err = migration.Up()
	if err != nil && err != migrate.ErrNoChange {
		return err
	}

	version, _, err := migration.Version()
	if err != nil {
		return err
	}

	migration.Log.Printf("Active database version: %d", version)

	return nil
}
