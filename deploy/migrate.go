package migrate

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database"
	"github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
)

// Init runs the migrations
func Init(db *sql.DB) (err error) {
	var (
		m      *migrate.Migrate
		driver database.Driver
	)

	if driver, err = pgx.WithInstance(db, &pgx.Config{}); err != nil {
		fmt.Printf("Could not create pgx driver instance: %v\n", err)
		return
	}

	if m, err = migrate.NewWithDatabaseInstance(
		"file://deploy/migrations",
		"pgx", driver); err != nil {
		fmt.Printf("Could not create migration instance: %v\n", err)
		return
	}

	m.Up()

	return
}
