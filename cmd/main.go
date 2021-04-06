package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/travis-james/RSSReader/readfeeds"
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

// For NHK.
func main() {
	a := readfeeds.GetNHK()
	fmt.Println(a)
}

// For news24
// func main() {
// 	//url := "https://www.news24.jp/rss/index.rdf"
// 	//resp, err := fetchFeed(url)
// 	resp, err := ioutil.ReadFile("ex.rdf")
// 	if err != nil {
// 		log.Fatalf("fetch feed failed: %v", err)
// 	}

// 	// Definitely wouldn't have found this out on my own:
// 	// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
// 	reader := bytes.NewReader(resp)
// 	decoder := xml.NewDecoder(reader)
// 	decoder.CharsetReader = charset.NewReaderLabel
// 	article := &Asahi{}
// 	err = decoder.Decode(article)

// 	//err = xml.Unmarshal(resp, article)
// 	if err != nil {
// 		log.Fatalf("xml unmarshal fail: %v", err)
// 	}
// 	fmt.Println(article)
// }

// For Asahi.
// func main() {
// 	//url := "https://www.news24.jp/rss/index.rdf"
// 	//resp, err := fetchFeed(url)
// 	resp, err := ioutil.ReadFile("ex.rdf")
// 	if err != nil {
// 		log.Fatalf("fetch feed failed: %v", err)
// 	}
// 	article := &Asahi{}
// 	err = xml.Unmarshal(resp, article)
// 	if err != nil {
// 		log.Fatalf("xml unmarshal fail: %v", err)
// 	}
// 	fmt.Println(article)
// }
