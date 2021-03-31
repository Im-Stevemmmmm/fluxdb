package database

import (
	"fmt"
	"time"
)

func NewLogData(msg string) *LogData {
	ts := time.Now().Format(time.RFC3339)
	return &LogData{
		Timestamp: ts,
		Message:   fmt.Sprintf("[%s] %s", ts, msg),
	}
}

// LogData is the data for logging an event to the logs.
type LogData struct {
	Timestamp string
	Message   string
}
