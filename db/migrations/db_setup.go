package migrations

import (
	"os"
	"social_media/db"

	"github.com/golang-migrate/migrate/v4"
	migratePostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func SetupDB() {
	config := migratePostgres.Config{}

	driver, err := migratePostgres.WithInstance(db.DataBase().DB, &config)
	if err != nil {
		panic(err)
	}

	migrationsDir := os.Getenv("MIGRATIONS_DIR")

	m, err := migrate.NewWithDatabaseInstance(
		"file://"+migrationsDir,
		"socialnetwork", driver)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		panic(err)
	}

}
