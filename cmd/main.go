package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
)

type News struct {
	NHK   []readfeeds.NHKItem
	Asahi []readfeeds.AsahiItem
	Twit  []mytwitter.TTrend
	Date  time.Time
	ID    int
}

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
)

// For NHK.
func main() {
	// Get NHK.
	n := &readfeeds.NHK{}
	err := readfeeds.SetFeed(n, nhkURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range n.XMLCh.Items {
	// 	fmt.Println(item)
	// }

	// Get Asahi.
	a := &readfeeds.Asahi{}
	err = readfeeds.SetFeed(a, asahiURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range a.Items {
	// 	fmt.Println(item)
	// }

	// Translate Asahi & NHK.
	readfeeds.TranslateTitle(n.XMLCh.Items, a.Items)
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

	todaysnews := News{
		NHK:   n.XMLCh.Items,
		Asahi: a.Items,
		Twit:  c,
		Date:  time.Now(),
		ID:    1,
	}
	fmt.Println(todaysnews)

	jnews, err := json.Marshal(todaysnews)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===\n===\n===\n====\n===\n===")
	fmt.Println(string(jnews))
}
