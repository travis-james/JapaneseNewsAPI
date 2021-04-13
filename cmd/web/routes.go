package main

import (
	"github.com/gorilla/mux"
)

// routes returns a mux with all the routes registered to it.
func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home).Methods("GET")
	router.HandleFunc("/update", app.updatenews).Methods("POST")
	router.HandleFunc("/{date}", app.getnews).Methods("GET")

	return router
}
