package utils_test

import (
	"testing"

	"data-ingestion/server/utils"

	"github.com/stretchr/testify/assert"
)

func TestValidateCPF(t *testing.T) {
	t.Run("should return true when CPF is valid", func(t *testing.T) {
		cpf := "042.098.288-40"
		res := utils.ValidateCPF(cpf)
		assert.True(t, res)
	})

	t.Run("should return false when CPF is invalid", func(t *testing.T) {
		cpf := "123.456.789-01"
		res := utils.ValidateCPF(cpf)
		assert.False(t, res)
	})
}

func TestValidateCNPJ(t *testing.T) {
	t.Run("should return true when CNPJ is valid", func(t *testing.T) {
		cnpj := "68.537.862/0001-91"
		res := utils.ValidateCNPJ(cnpj)
		assert.True(t, res)
	})

	t.Run("should return true when CNPJ is valid and first digit is 0", func(t *testing.T) {
		cnpj := "42.890.578/0001-00"
		res := utils.ValidateCNPJ(cnpj)
		assert.True(t, res)
	})

	t.Run("should return false when CNPJ is invalid", func(t *testing.T) {
		cnpj := "68.537.862/0001-92"
		res := utils.ValidateCNPJ(cnpj)
		assert.False(t, res)
	})
}
