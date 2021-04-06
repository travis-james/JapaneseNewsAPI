package readfeeds

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"log"

	"golang.org/x/net/html/charset"
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
func GetNHK() *NHK {
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

type Asahi struct {
	XMLName xml.Name    `xml:"RDF"`
	Items   []AsahiItem `xml:"item"`
}

type AsahiItem struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"dc:date"`
}

// For Asahi.
func GetAsahi() *Asahi {
	//url := "https://www.news24.jp/rss/index.rdf"
	//resp, err := fetchFeed(url)
	resp, err := ioutil.ReadFile("ex.rdf")
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

// For news24
func GetNews24() *Asahi {
	//url := "https://www.news24.jp/rss/index.rdf"
	//resp, err := fetchFeed(url)
	resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		log.Fatalf("fetch feed failed: %v", err)
	}

	// Definitely wouldn't have found this out on my own:
	// https://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
	reader := bytes.NewReader(resp)
	decoder := xml.NewDecoder(reader)
	decoder.CharsetReader = charset.NewReaderLabel
	article := &Asahi{}
	err = decoder.Decode(article)

	//err = xml.Unmarshal(resp, article)
	if err != nil {
		log.Fatalf("xml unmarshal fail: %v", err)
	}
	return article
}
