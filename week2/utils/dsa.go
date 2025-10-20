// Package utils provides a set of utility as well as data structures that may be of use
// in the entire application.
package utils

import (
	"database/sql/driver"
	"errors"
	"strings"
)

type StringSet []string

func (s *StringSet) Scan(value any) error {
	if value == nil {
		*s = nil
		return nil
	}

	// The database returns a []byte (which can be cast to string)
	// that contains a comma-separated list of values.
	// Courtesy of Gemini
	str, ok := value.([]byte)
	if !ok {
		// Handle case where value is already a string
		if sVal, isStr := value.(string); isStr {
			str = []byte(sVal)
		} else {
			// If it's not a byte slice or string, return an error
			return errors.New("unsupported type for StringSet scanning")
		}
	}

	// Split the comma-separated string into a slice of strings
	if len(str) == 0 {
		*s = nil
	} else {
		// Use the strings package to split the value
		*s = strings.Split(string(str), ",")
	}
	return nil
}

// Value implements the driver.Valuer interface.
// It converts the Go []string back into a comma-separated string for the database.
func (s StringSet) Value() (driver.Value, error) {
	if len(s) == 0 {
		return nil, nil
	}

	// Join the slice of strings back into a comma-separated string
	return strings.Join(s, ","), nil
}
