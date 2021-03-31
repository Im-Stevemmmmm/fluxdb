package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Im-Stevemmmmm/fluxdb/database"
)

var (
	insertMethods = []httpMethod{
		GET,
		POST,
	}
	insertEndpoint = endpoint{
		Path:        "/_insert",
		HandlerFunc: insertHandler,
		Methods:     insertMethods,
	}
)

func insertHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if errObj := validateRequest(r, insertMethods); errObj != nil {
		json.NewEncoder(w).Encode(errObj)
		return
	}

	var body insertRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		err := NewRequestError(RuntimeError, "Invalid body")
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := body.isValid(); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	err = database.LightDB.Set(*body.Key, *body.Value)
	if err != nil {
		NewRequestError(RuntimeError, "Unable to set value")
	}

	duration := time.Since(startTime)

	json.NewEncoder(w).Encode(insertResult{
		InsertCount: 1,
		InsertTime:  duration.Milliseconds(),
	})
}

func (i insertRequest) isValid() *RequestError {
	if i.Key == nil {
		return NewRequestError(MalformedRequest, "Missing key")
	}
	if i.Value == nil {
		return NewRequestError(MalformedRequest, "Missing value")
	}
	return nil
}

type insertRequest struct {
	Key   *string      `json:"key"`
	Value *interface{} `json:"value"`
}

type insertResult struct {
	InsertCount int    `json:"insert_count"`
	InsertTime  int64  `json:"insert_time"`
	Error       string `json:"error,omitempty"`
}
