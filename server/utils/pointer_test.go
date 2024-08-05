package utils_test

import (
	"testing"
	"time"

	"data-ingestion/server/utils"

	"github.com/stretchr/testify/assert"
)

func TestNewPointerBool(t *testing.T) {
	t.Run("should return a pointer to a bool", func(t *testing.T) {
		value := true
		res := utils.NewPointerBool(value)
		assert.Equal(t, &value, res)
	})
}

func TestNewPointerFloat64(t *testing.T) {
	t.Run("should return a pointer to a float64", func(t *testing.T) {
		value := 100.0
		res := utils.NewPointerFloat64(value)
		assert.Equal(t, &value, res)
	})
}

func TestNewPointerString(t *testing.T) {
	t.Run("should return a pointer to a string", func(t *testing.T) {
		value := "value"
		res := utils.NewPointerString(value)
		assert.Equal(t, &value, res)
	})
}

func TestNewPointerTime(t *testing.T) {
	t.Run("should return a pointer to a time.Time", func(t *testing.T) {
		value := "2020-01-01"
		date, _ := time.Parse("2006-01-02", value)
		res := utils.NewPointerTime(date)
		assert.Equal(t, &date, res)
	})
}
