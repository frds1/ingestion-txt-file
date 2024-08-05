package ingestion

import "context"

// IIngestion is the interface of the ingestion domain
type IIngestion interface {
	CreateDataIngestion(ctx context.Context, ingestion []*Ingestion) (err error)
}
