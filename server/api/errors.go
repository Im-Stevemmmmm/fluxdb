package api

import (
	"encoding/json"
	"net/http"
	"time"
)

type restError string

const (
	MalformedRequest   restError = "MALFORMED_REQUEST"
	ImproperHttpMethod           = "IMPROPER_HTTP_METHOD"
	RuntimeError                 = "RUNTIME_ERROR"
)

func (r RequestError) Write(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(r)
}

// NewRequestError creates a new APIRequestError and initializes the time
// to the current time.
func NewRequestError(error restError, message string) *RequestError {
	return &RequestError{
		Time:    time.Now().Format(time.RFC3339),
		Error:   error,
		Message: message,
	}
}

// APIRequestError is an error produced by a REST api request.
type RequestError struct {
	Time    string    `json:"time"`
	Error   restError `json:"error"`
	Message string    `json:"message"`
}
