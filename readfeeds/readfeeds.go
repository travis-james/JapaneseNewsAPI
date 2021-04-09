package readfeeds

import (
	"encoding/json"
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

// GetAsahi goes to the Asahi RSS feed and turns that xml response into a struct. That struct is then marshaled into
// a JSON string that is returned. The JSON string contains articles of the form [{Title, Link, Date}, etc].
func GetNHK() string {
	// First fetch the RSS feed.
	resp, err := fetchFeed(nhkURL)
	// resp, err := ioutil.ReadFile("ex.xml")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}

	// Turn the xml response into a struct.
	feed := &NHK{}
	err = xml.Unmarshal(resp, feed)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}

	// Now turn the struct's items into a json []byte.
	article, err := json.Marshal(feed.XMLCh.Items)
	if err != nil {
		log.Fatalf("json marshal fail: %v", err)
	}
	return string(article)
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

// GetAsahi goes to the Asahi RSS feed and turns that xml response into a struct. That struct is then marshaled into
// a JSON string that is returned. The JSON string contains articles of the form [{Title, Link, Date}, etc].
func GetAsahi() string {
	// First fetch the RSS feed.
	resp, err := fetchFeed(asahiURL)
	// resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}

	// Turn the xml response into a struct.
	feed := &Asahi{}
	err = xml.Unmarshal(resp, feed)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}

	// Now turn the struct's items into a json []byte.
	article, err := json.Marshal(feed.Items)
	if err != nil {
		log.Fatalf("json marshal fail: %v", err)
	}
	return string(article)
}

func fetchFeed(url string) ([]byte, error) {
	net := &http.Client{
		Timeout: time.Second * 10, // Having timeout is a good practice, I need to remember to do this.
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
