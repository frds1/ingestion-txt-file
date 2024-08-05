package utils_test

import (
	"testing"

	"data-ingestion/server/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseBool(t *testing.T) {
	t.Run("should return true when value is 1", func(t *testing.T) {
		value := "1"
		expected := true
		res := utils.ParseBool(value)
		assert.Equal(t, expected, *res)
	})

	t.Run("should return false when value is 0", func(t *testing.T) {
		value := "0"
		expected := false
		res := utils.ParseBool(value)
		assert.Equal(t, expected, *res)
	})

	t.Run("should return nil when value is not 0 or 1", func(t *testing.T) {
		value := "2"
		res := utils.ParseBool(value)
		require.Nil(t, res)
	})
}

func TestParseFloat(t *testing.T) {
	t.Run("should return a float when value is a number", func(t *testing.T) {
		value := "100.0"
		expected := 100.0
		res := utils.ParseFloat(value)
		assert.Equal(t, expected, *res)
	})

	t.Run("should return nil when value is not a number", func(t *testing.T) {
		value := "a"
		res := utils.ParseFloat(value)
		require.Nil(t, res)
	})
}

func TestParseString(t *testing.T) {
	t.Run("should return a string when value is not NULL", func(t *testing.T) {
		value := "value"
		expected := "value"
		res := utils.ParseString(value)
		assert.Equal(t, expected, *res)
	})

	t.Run("should return nil when value is NULL", func(t *testing.T) {
		value := "NULL"
		res := utils.ParseString(value)
		require.Nil(t, res)
	})
}

func TestParseDate(t *testing.T) {
	t.Run("should return a date when value is a date", func(t *testing.T) {
		value := "2020-01-01"
		expected := "2020-01-01"
		res := utils.ParseDate(value)
		assert.Equal(t, expected, res.Format("2006-01-02"))
	})

	t.Run("should return nil when value is not a date", func(t *testing.T) {
		value := "a"
		res := utils.ParseDate(value)
		require.Nil(t, res)
	})
}
