// Package utils provides a set of random utility functions and handlings.
package utils

import (
	"encoding/json"
	"io"
	"os"
)

var logger io.Writer

func InitLogger() {
	f, _ := os.OpenFile("/var/log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logger = f
}

// LogJSON sends a log message with json after marshalling.
// This is a Fire-and-forget function.
func LogJSON(v any) {
	go func(data any) {
		b, _ := json.Marshal(data)
		logger.Write(append(b, '\n'))
	}(v)
}
