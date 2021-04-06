package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/travis-james/RSSReader/readfeeds"
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

// For NHK.
func main() {
	a := readfeeds.GetNHK()
	fmt.Println(a)
	// b := readfeeds.GetAsahi()
	// fmt.Println(b)
}
