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

	if err := validateRequest(r, insertMethods); err != nil {
		err.Write(w)
		return
	}

	var body insertRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		NewRequestError(RuntimeError, "Invalid body").Write(w)
		return
	}

	if err := body.IsValid(); err != nil {
		err.Write(w)
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

func (i insertRequest) IsValid() *RequestError {
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
