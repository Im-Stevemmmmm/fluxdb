package api

import (
	"fmt"
	"net/http"
	"strings"
)

var (
	// Endpoints is the available endpoints for performing actions on the database
	Endpoints = []endpoint{
		insertEndpoint,
	}
)

type endpoint struct {
	Path        string
	HandlerFunc func(http.ResponseWriter, *http.Request)
	Methods     []httpMethod
}

type httpMethod string

const (
	GET    httpMethod = http.MethodGet
	POST              = http.MethodPost
	PUT               = http.MethodPut
	DELETE            = http.MethodDelete
)

func validateRequest(r *http.Request, methods []httpMethod) *RequestError {
	m := r.Method

	var contains bool
	for _, v := range methods {
		if httpMethod(m) == v {
			contains = true
			break
		}
	}

	if !contains {
		elems := make([]string, len(methods))
		for i, v := range methods {
			elems[i] = string(v)
		}

		m := strings.Join(elems, " or ")
		msg := fmt.Sprintf("Http method must be %s", m)

		return NewRequestError(ImproperHttpMethod, msg)
	}

	return nil
}
