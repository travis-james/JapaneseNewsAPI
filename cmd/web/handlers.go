package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
)

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
	prevDate = ""
)

// home (GET) just writes a small welcome message explaining what the api does.
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	info := []byte(`Welcome to JapaneseNewsAPI. This is a web api that stores the headlines 
from the RSS feeds of NHK and Asahi news, as well as what was trending on
Twitter in Japan at that time.

/updatenews 'POST' => update the database with the latest news.
/getnews/1999-12-31 'GET' => get news for that given date. Response is JSON.`)
	w.Write(info)
}

// insertNews (POST) will get the latest news/trends and store it in the database,
//  if it isn't already in the database.
func (app *application) insertNews(w http.ResponseWriter, r *http.Request) {
	// First check that the date isn't already in the database.
	todaysDate := time.Now().Format("2006-01-02")
	retval, _ := app.news.Get(todaysDate)
	if retval != nil {
		msg := todaysDate + " already exists in the database, no action taken.\n"
		w.Write([]byte(msg))
		return
	}

	// Get NHK.
	n := &mynews.NHK{}
	err := mynews.SetFeed(n, nhkURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get Asahi.
	a := &mynews.Asahi{}
	err = mynews.SetFeed(a, asahiURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Translate Asahi & NHK.
	mynews.TranslateTitle(n.XMLCh.Items, a.Items)

	// Get twitter trends.
	c, err := mytwitter.GetTrends()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// For DB.
	_, err = app.news.Insert(*n, *a, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	msg := "Succesfully inserted the news of the following date: " + todaysDate + "\n"
	w.Write([]byte(msg))
}

// getNews (GET) retrieves a news struct based on date (if it exists) and returns
// a JSON object.
func (app *application) getNews(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")

	vars := mux.Vars(r)
	d, exists := vars["date"]
	if exists == false {
		http.Error(w, errors.New("getnews: bad date").Error(), http.StatusBadRequest)
		return
	}
	// Find in db.
	urnews, err := app.news.Get(d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(urnews)
	return
}
