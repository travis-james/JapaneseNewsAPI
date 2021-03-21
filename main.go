package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type Asahi struct {
	XMLName xml.Name    `xml:"RDF"`
	Items   []AsahiItem `xml:"item"`
}

type AsahiItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"dc:date"`
}

func fetchFeed(url string) ([]byte, error) {
	net := &http.Client{
		Timeout: time.Second * 10,
	}

	// Get the feed.
	res, err := net.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func main() {
	// url := "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	// resp, err := fetchFeed(url)
	resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}
	article := &Asahi{}
	err = xml.Unmarshal(resp, article)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}
	fmt.Println(article)
}
