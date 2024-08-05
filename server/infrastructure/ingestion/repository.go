package ingestion

import (
	"context"
	ingestionDomain "data-ingestion/server/domain/ingestion"
	"database/sql"
)

type IngestionRepository struct {
	DB *sql.DB
}

// CreateDataIngestion creates a new data ingestion
func (r *IngestionRepository) CreateDataIngestion(ctx context.Context, ingestion []*ingestionDomain.Ingestion) (err error) {
	pg := &PGIngestion{DB: r.DB}
	return pg.CreateDataIngestion(ctx, ingestion)
}
