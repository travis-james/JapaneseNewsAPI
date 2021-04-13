package main

import (
	"net/http"
)

func (app *application) updatenews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from updatenews"))
}

func (app *application) getnews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from getnews"))
}
