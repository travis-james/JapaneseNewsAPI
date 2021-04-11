package main

import (
	"fmt"
	"log"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
)

type News struct {
	NHK   []readfeeds.NHKItem
	Asahi []readfeeds.AsahiItem
	Twit  *mytwitter.TTrends
	Date  time.Time
}

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
)

// For NHK.
func main() {
	n := &readfeeds.NHK{}
	err := readfeeds.SetFeed(n, nhkURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range n.XMLCh.Items {
	// 	fmt.Println(item)
	// }

	a := &readfeeds.Asahi{}
	err = readfeeds.SetFeed(a, asahiURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	for _, item := range a.Items {
		fmt.Println(item)
	}
	// c, err := mytwitter.GetTrends()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, trend := range c.Trends {
	// 	fmt.Println(trend)
	// }

}
