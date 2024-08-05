package ingestion_test

import (
	"context"
	"testing"
	"time"

	ingestionDomain "data-ingestion/server/domain/ingestion"
	ingestionRepo "data-ingestion/server/infrastructure/ingestion"
	"data-ingestion/server/utils"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateDataIngestion(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	lastPurchaseDate, _ := time.Parse("2024-01-02", "2024-01-01")

	ingestion := &ingestionDomain.Ingestion{
		CPF:                utils.NewPointerString("042.098.288-40"),
		Private:            utils.NewPointerBool(true),
		Incomplete:         utils.NewPointerBool(false),
		LastPurchaseDate:   utils.NewPointerTime(lastPurchaseDate),
		AverageTicket:      utils.NewPointerFloat64(100.0),
		LastPurchaseTicket: utils.NewPointerFloat64(200.0),
		MostFrequentStore:  utils.NewPointerString("68.537.862/0001-91"),
		LastPurchaseStore:  utils.NewPointerString("68.537.862/0001-91"),
	}

	repo := &ingestionRepo.IngestionRepository{
		DB: db,
	}

	t.Run("should create data ingestion", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO ingestions \(cpf, private, incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
			WithArgs(
				ingestion.CPF,
				ingestion.Private,
				ingestion.Incomplete,
				ingestion.LastPurchaseDate,
				ingestion.AverageTicket,
				ingestion.LastPurchaseTicket,
				ingestion.MostFrequentStore,
				ingestion.LastPurchaseStore).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit()

		err = repo.CreateDataIngestion(context.Background(), []*ingestionDomain.Ingestion{ingestion})
		assert.NoError(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to begin transaction", func(t *testing.T) {
		mock.ExpectBegin().WillReturnError(assert.AnError)

		err = repo.CreateDataIngestion(context.Background(), []*ingestionDomain.Ingestion{ingestion})
		assert.Error(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to execute query", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO ingestions \(cpf, private, incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
			WillReturnError(assert.AnError)

		mock.ExpectRollback()

		err = repo.CreateDataIngestion(context.Background(), []*ingestionDomain.Ingestion{ingestion})
		assert.Error(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("should return error when failed to commit transaction", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO ingestions \(cpf, private, incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store\) VALUES \(\$1, \$2, \$3, \$4, \$5, \$6, \$7, \$8\)`).
			WillReturnResult(sqlmock.NewResult(1, 1))

		mock.ExpectCommit().WillReturnError(assert.AnError)

		err = repo.CreateDataIngestion(context.Background(), []*ingestionDomain.Ingestion{ingestion})
		assert.Error(t, err)

		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})
}
