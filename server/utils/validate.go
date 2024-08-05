package utils

import (
	"strconv"
	"strings"
)

const (
	// CPFSize is the size of a CPF
	CPFSize = 11
	// CNPJSize is the size of a CNPJ
	CNPJSize = 14

	// CPFFirstCheckDigitInitialPosition  is the initial position of the first check digit of a CPF
	CPFFirstCheckDigitInitialPosition = 10
	// CPFSecondCheckDigitInitialPosition  is the initial position of the second check digit of a CPF
	CPFSecondCheckDigitInitialPosition = 11

	// CNPJFirstCheckDigitInitialPosition is the initial position of the first check digit of a CNPJ
	CNPJFirstCheckDigitInitialPosition = 5
	// CNPJSecondCheckDigitInitialPosition is the initial position of the second check digit of a CNPJ
	CNPJSecondCheckDigitInitialPosition = 6

	// CNPJFirstCheckDigitPosition is the position of the first check digit of a CNPJ
	CNPJFirstCheckDigitPosition = 12
	// CNPJSecondCheckDigitPosition is the position of the second check digit of a CNPJ
	CNPJSecondCheckDigitPosition = 13
)

// ValidateCPF validates a CPF
func ValidateCPF(cpf string) bool {
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	return len(cpf) == CPFSize && ValidateCPFCheckDigit(cpf, CPFFirstCheckDigitInitialPosition) && ValidateCPFCheckDigit(cpf, CPFSecondCheckDigitInitialPosition)
}

// ValidateCPFCheckDigit validates a CPF check digit
func ValidateCPFCheckDigit(cpf string, initialPosition int) bool {
	sum := 0
	for i := 0; i < initialPosition-1; i++ {
		digit, _ := strconv.Atoi(string(cpf[i]))
		sum += digit * (initialPosition - i)
	}

	rest := (sum * 10) % 11
	if rest == 10 {
		rest = 0
	}

	CheckDigit, _ := strconv.Atoi(string(cpf[initialPosition-1]))
	return CheckDigit == rest
}

// ValidateCNPJ validates a CNPJ
func ValidateCNPJ(cnpj string) bool {
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")

	return len(cnpj) == CNPJSize && ValidateCNPJCheckDigit(cnpj, CNPJFirstCheckDigitPosition, CNPJFirstCheckDigitInitialPosition) && ValidateCNPJCheckDigit(cnpj, CNPJSecondCheckDigitPosition, CNPJSecondCheckDigitInitialPosition)
}

// ValidateCNPJCheckDigit validates a CNPJ check digit
func ValidateCNPJCheckDigit(cnpj string, pos int, initialPosition int) bool {
	sum := 0
	weight := initialPosition

	for i := 0; i < pos; i++ {
		digit, _ := strconv.Atoi(string(cnpj[i]))
		sum += digit * weight
		weight--
		if weight < 2 {
			weight = 9
		}
	}

	remainder := sum % 11
	if remainder < 2 {
		remainder = 0
	} else {
		remainder = 11 - remainder
	}

	checkDigit, _ := strconv.Atoi(string(cnpj[pos]))
	return checkDigit == remainder
}
