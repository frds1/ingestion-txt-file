package utils

import (
	"strconv"
	"strings"
	"time"
)

// ParseBool receives a string and returns a pointer to a bool
func ParseBool(value string) *bool {
	switch value {
	case "0":
		falseVal := false
		return &falseVal
	case "1":
		trueVal := true
		return &trueVal
	default:
		return nil
	}
}

// ParseFloat receives a string and returns a pointer to a float64
func ParseFloat(value string) *float64 {
	value = strings.Replace(value, ",", ".", -1)
	if floatVal, err := strconv.ParseFloat(value, 64); err == nil {
		return &floatVal
	}
	return nil
}

// ParseString receives a string and returns a pointer to a string
func ParseString(value string) *string {
	if value == "NULL" {
		return nil
	}
	return &value
}

// ParseDate receives a string and returns a pointer to a time.Time
func ParseDate(value string) *time.Time {
	date, err := time.Parse("2006-01-02", value)
	if err == nil {
		return &date
	}
	return nil
}
