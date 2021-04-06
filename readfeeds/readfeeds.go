package readfeeds

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

var URLs = []string{
	"http://www.asahi.com/rss/asahi/newsheadlines.rdf",
	"https://www.news24.jp/rss/index.rdf",
	"https://www.nhk.or.jp/rss/news/cat0.xml",
}

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
func getNHK() *NHK {
	// url := "https://www.news24.jp/rss/index.rdf"
	// resp, err := fetchFeed(url)
	resp, err := ioutil.ReadFile("ex.xml")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}
	article := &NHK{}
	// reader := bytes.NewReader(resp)
	// decoder := xml.NewDecoder(reader)
	// decoder.CharsetReader = charset.NewReaderLabel
	// err = decoder.Decode(article)
	err = xml.Unmarshal(resp, article)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}
	return article
}
