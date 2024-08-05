package utils

import "time"

// NewPointerBool returns a pointer to a bool
func NewPointerBool(value bool) *bool {
	return &value
}

// NewPointerFloat64 returns a pointer to a float64
func NewPointerFloat64(value float64) *float64 {
	return &value
}

// NewPointerString returns a pointer to a string
func NewPointerString(value string) *string {
	return &value
}

// NewPointerTime returns a pointer to a time.Time
func NewPointerTime(value time.Time) *time.Time {
	return &value
}
