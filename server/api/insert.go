package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Im-Stevemmmmm/fluxdb/database/document"
)

var (
	insertMethods = []httpMethod{
		GET,
		POST,
	}
	insertEndpoint = endpoint{
		Path:        "/{namespace}/{index}/_field",
		HandlerFunc: insertHandler,
		Methods:     insertMethods,
	}
)

func insertHandler(w http.ResponseWriter, r *http.Request) {
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

	vars := mux.Vars(r)
	ns := vars["namespace"]
	index := vars["index"]

	res, err := document.Insert(&document.InsertCommandInput{
		Namespace: ns,
		Index:     index,
		Payload:   body.Document,
	})
	json.NewEncoder(w).Encode(res)
}

func (i insertRequest) isValid() *RequestError {
	if i.Document == nil {
		return NewRequestError(MalformedRequest, "Missing insertion data")
	}
	return nil
}

type insertRequest struct {
	Document         document.Document `json:"document"`
	ReturnAttributes []string          `json:"return_attributes"`
}

type insertResult struct {
	InsertCount int    `json:"insert_count"`
	InsertTime  int64  `json:"insert_time"`
	Error       string `json:"error,omitempty"`
}
