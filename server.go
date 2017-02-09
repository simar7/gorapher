package main

import "github.com/gorilla/mux"

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/stats/{type}", StatsHandler)
	return r
}
