package ingestion

import "time"

const (
	// Indexes for the fields in the file
	CPFIndex                = 0
	PrivateIndex            = 1
	IncompleteIndex         = 2
	LastPurchaseDateIndex   = 3
	AverageTicketIndex      = 4
	LastPurchaseTicketIndex = 5
	MostFrequentStoreIndex  = 6
	LastPurchaseStoreIndex  = 7
)

// Ingestion represents the ingestion entity
type Ingestion struct {
	CPF                *string
	Private            *bool
	Incomplete         *bool
	LastPurchaseDate   *time.Time
	AverageTicket      *float64
	LastPurchaseTicket *float64
	MostFrequentStore  *string
	LastPurchaseStore  *string
}
