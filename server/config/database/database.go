package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// OpenConnection open and test connection with database
func OpenConnection() (database *sql.DB, err error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=public", os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME"))

	if database, err = sql.Open("pgx", connectionString); err != nil {
		return
	}

	if err = database.Ping(); err != nil {
		return
	}

	return
}
