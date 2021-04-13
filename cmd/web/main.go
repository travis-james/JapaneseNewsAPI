package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
	"go.mongodb.org/mongo-driver/mongo"
)

type News struct {
	NHK   []readfeeds.NHKItem
	Asahi []readfeeds.AsahiItem
	Twit  []mytwitter.TTrend
	Date  string
	ID    int
}

// application struct allows dependency injection. For example,
// by using this, my handlers could have access to using
//  logs I declare in main. As of now it's just for DB.
type application struct {
}

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
	client   *mongo.Client
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %v\n", *addr)
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
