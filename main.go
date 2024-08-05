package main

import (
	"context"
	"database/sql"
	"log"
	"os"

	deploy "data-ingestion/deploy"
	app "data-ingestion/server/application/ingestion"
	database "data-ingestion/server/config/database"

	"github.com/joho/godotenv"
)

func main() {
	var (
		err  error
		file *os.File
		db   *sql.DB
		ctx  = context.Background()
	)

	if err = godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	if db, err = database.OpenConnection(); err != nil {
		log.Fatal("Could not create postgres connection: ", err)
	}

	if err = deploy.Init(db); err != nil {
		log.Fatal("Could not migrate database: ", err)
	}

	ingestionService := app.NewService(db)

	if file, err = os.Open("Base.txt -Teste Técnico .txt"); err != nil {
		log.Fatal("Could not open file: ", err)
	}

	if err = ingestionService.CreateDataIngestion(ctx, file); err != nil {
		log.Fatal("Could not create data ingestion: ", err)
	}

}
