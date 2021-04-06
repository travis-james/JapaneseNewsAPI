package readfeeds

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	asahiURL = "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	nhkURL   = "https://www.nhk.or.jp/rss/news/cat0.xml"
)

type NHK struct {
	XMLName xml.Name   `xml:"rss"`
	XMLCh   NHKChannel `xml:"channel"`
}

type NHKChannel struct {
	Items []NHKItem `xml:"item"`
}

type NHKItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"pubDate"`
}

// For NHK.
func GetNHK() *NHK {
	resp, err := fetchFeed(nhkURL)
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}
	article := &NHK{}
	err = xml.Unmarshal(resp, article)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}
	return article
}

type Asahi struct {
	XMLName xml.Name    `xml:"RDF"`
	Items   []AsahiItem `xml:"item"`
}

type AsahiItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"date"`
}

// For Asahi.
func GetAsahi() *Asahi {
	url := "http://www.asahi.com/rss/asahi/newsheadlines.rdf"
	resp, err := fetchFeed(url)
	//resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}
	article := &Asahi{}
	err = xml.Unmarshal(resp, article)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}
	return article
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
