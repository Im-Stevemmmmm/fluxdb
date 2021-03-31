package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Im-Stevemmmmm/fluxdb/database"
)

var (
	readMethods = []httpMethod{
		GET,
		POST,
	}
	readEndpoint = endpoint{
		Path:        "/_read",
		HandlerFunc: readHandler,
		Methods:     readMethods,
	}
)

func readHandler(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()

	if err := validateRequest(r, insertMethods); err != nil {
		err.Write(w)
		return
	}

	var body readRequest

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		NewRequestError(RuntimeError, "Invalid body").Write(w)
		return
	}

	if err := body.IsValid(); err != nil {
		err.Write(w)
		return
	}

	value, err := database.Instance.Get(*body.Key)
	if err != nil {
		NewRequestError(RuntimeError, "Could not read key").Write(w)
		return
	}

	duration := time.Since(startTime)

	json.NewEncoder(w).Encode(readResult{
		Data:     value,
		ReadTime: duration.Milliseconds(),
	})
}

func (r readRequest) IsValid() *RequestError {
	if r.Key == nil {
		return NewRequestError(MalformedRequest, "Missing key")
	}
	return nil
}

type readRequest struct {
	Key *string
}

type readResult struct {
	Data     interface{} `json:"data"`
	ReadTime int64       `json:"read_time"`
}
