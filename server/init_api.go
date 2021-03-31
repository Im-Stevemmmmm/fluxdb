package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Im-Stevemmmmm/fluxdb/api"
)

func initAPI() {
	r := mux.NewRouter()
	for _, e := range api.Endpoints {
		r.HandleFunc(e.Path, e.HandlerFunc)
	}
	http.Handle("/", r)
}
