package main

import (
	"github.com/gorilla/mux"
)

// routes returns a mux with all the routes registered to it.
func (app *application) routes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home).Methods("GET")
	router.HandleFunc("/update", app.insertNews).Methods("POST") // Maybe not post?
	router.HandleFunc("/getnews/{date}", app.getNews).Methods("GET")

	return router
}
