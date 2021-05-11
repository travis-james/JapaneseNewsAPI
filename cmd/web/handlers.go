package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/pkg/models"
)

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
	id       = 0
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	info := []byte(`Welcome to JapaneseNewsAPI. This is a web api that stores the headlines 
from the RSS feeds of NHK and Asahi news, as well as what was trending on
Twitter in Japan at that time.

/updatenews 'POST' => update the database with the latest news.
/getnews 'GET' => get news for that given date. Response is JSON.`)
	w.Write(info)
}

func (app *application) updatenews(w http.ResponseWriter, r *http.Request) {
	// Get NHK.
	n := &mynews.NHK{}
	err := mynews.SetFeed(n, nhkURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range n.XMLCh.Items {
	// 	fmt.Println(item)
	// }

	// Get Asahi.
	a := &mynews.Asahi{}
	err = mynews.SetFeed(a, asahiURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range a.Items {
	// 	fmt.Println(item)
	// }

	// Translate Asahi & NHK.
	mynews.TranslateTitle(n.XMLCh.Items, a.Items)
	// for i, item := range a.Items {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(item)
	// }
	// for i, item := range n.XMLCh.Items {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(item)
	// }

	// Get twitter trends.
	c, err := mytwitter.GetTrends()
	if err != nil {
		log.Fatalln(err)
	}
	// for _, trend := range c.Trends {
	// 	fmt.Println(trend)
	// }

	date := time.Now().Format("2006-01-02")
	todaysnews := models.News{
		NHK:   n.XMLCh.Items,
		Asahi: a.Items,
		Twit:  c,
		Date:  date,
		ID:    id,
	}
	id++
	//fmt.Println(todaysnews)

	// JSON?
	// jnews, err := json.Marshal(todaysnews)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(jnews))

	// For DB.
	collection := app.client.Database("jpnews").Collection("day")
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(app.ctx, todaysnews)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func (app *application) getnews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from getnews"))
}
