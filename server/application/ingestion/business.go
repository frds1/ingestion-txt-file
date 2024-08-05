package ingestion

import (
	"bufio"
	"context"
	"database/sql"
	"os"
	"strings"

	domain "data-ingestion/server/domain/ingestion"
	ingestionRepo "data-ingestion/server/infrastructure/ingestion"
	"data-ingestion/server/utils"
)

type Service struct {
	DB domain.IIngestion
}

func NewService(db *sql.DB) *Service {
	return &Service{
		DB: &ingestionRepo.IngestionRepository{
			DB: db,
		},
	}
}

// CreateDataIngestion creates a new data ingestion
func (s *Service) CreateDataIngestion(ctx context.Context, file *os.File) (err error) {
	var (
		scanner    *bufio.Scanner
		ingestions []*domain.Ingestion
	)

	if scanner = bufio.NewScanner(file); !scanner.Scan() {
		return
	}

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		ingestions = append(ingestions, parseData(fields))
	}

	if err = scanner.Err(); err != nil {
		return
	}

	if err = s.DB.CreateDataIngestion(ctx, ingestions); err != nil {
		return
	}

	return
}

// parseData parses the data from the file
func parseData(fields []string) (ingestion *domain.Ingestion) {
	ingestion = &domain.Ingestion{
		CPF:                utils.ParseString(fields[domain.CPFIndex]),
		Private:            utils.ParseBool(fields[domain.PrivateIndex]),
		Incomplete:         utils.ParseBool(fields[domain.IncompleteIndex]),
		LastPurchaseDate:   utils.ParseDate(fields[domain.LastPurchaseDateIndex]),
		AverageTicket:      utils.ParseFloat(fields[domain.AverageTicketIndex]),
		LastPurchaseTicket: utils.ParseFloat(fields[domain.LastPurchaseTicketIndex]),
		MostFrequentStore:  utils.ParseString(fields[domain.MostFrequentStoreIndex]),
		LastPurchaseStore:  utils.ParseString(fields[domain.LastPurchaseStoreIndex]),
	}

	if validateFields(ingestion) {
		return &domain.Ingestion{}
	}

	return
}

// validateFields validates the fields of the ingestion
func validateFields(ingestion *domain.Ingestion) bool {
	return !validateField(ingestion.CPF, utils.ValidateCPF) ||
		!validateField(ingestion.LastPurchaseStore, utils.ValidateCNPJ) ||
		!validateField(ingestion.MostFrequentStore, utils.ValidateCNPJ)
}

// validateField validates a field
func validateField(field *string, validator func(string) bool) bool {
	if field != nil {
		return validator(*field)
	}

	return true
}
