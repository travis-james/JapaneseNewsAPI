package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// routes returns a mux with all the routes registered to it.
func (app *application) routes() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/", app.home).Methods("GET")
	router.HandleFunc("/updatenews", app.insertNews).Methods("POST") // Maybe not post?
	router.HandleFunc("/getnews/{date}", app.getNews).Methods("GET")

	return app.rateLimit(router)
}
