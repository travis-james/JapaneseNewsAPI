package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

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
	url := "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	resp, err := fetchFeed(url)
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}
	fmt.Println(string(resp))
}
