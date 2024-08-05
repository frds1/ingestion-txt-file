package ingestion

import (
	"context"
	ingestionDomain "data-ingestion/server/domain/ingestion"
	"database/sql"
)

type PGIngestion struct {
	DB *sql.DB
}

// CreateDataIngestion creates a new data ingestion
func (PG *PGIngestion) CreateDataIngestion(ctx context.Context, ingestions []*ingestionDomain.Ingestion) (err error) {
	var tx *sql.Tx

	if tx, err = PG.DB.Begin(); err != nil {
		return err
	}
	defer tx.Rollback()

	for _, ingestion := range ingestions {
		if _, err = tx.ExecContext(ctx,
			`INSERT INTO ingestions (cpf, private, incomplete, last_purchase_date, average_ticket, last_purchase_ticket, most_frequent_store, last_purchase_store)
		 VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			ingestion.CPF,
			ingestion.Private,
			ingestion.Incomplete,
			ingestion.LastPurchaseDate,
			ingestion.AverageTicket,
			ingestion.LastPurchaseTicket,
			ingestion.MostFrequentStore,
			ingestion.LastPurchaseStore,
		); err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return
}
