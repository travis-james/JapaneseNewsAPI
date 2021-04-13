package mynews

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytranslate"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
)

type News struct {
	NHK   []NHKItem
	Asahi []AsahiItem
	Twit  []mytwitter.TTrend
	Date  string
	ID    int
}

// NHK is the top level struct of reading the NHK RSS feed.
// XMLCh contains the point of interest, NHKChannel struct.
type NHK struct {
	XMLName xml.Name   `xml:"rss"`
	XMLCh   NHKChannel `xml:"channel"`
}

// NHKChannel contains a slice of NHKItems (headlines/articles).
type NHKChannel struct {
	Items []NHKItem `xml:"item"`
}

// NHKItem is the fundamental NHK data type containing a headline, link, date, etc.
type NHKItem struct {
	Title   string `xml:"title"`
	TitleEN string
	Link    string `xml:"link"`
	Date    string `xml:"pubDate"`
}

// Asahi is the top level struct of reading the Asahi RSS feed.
// Items contains the point of interest, AsahiItem struct.
type Asahi struct {
	XMLName xml.Name    `xml:"RDF"`
	Items   []AsahiItem `xml:"item"`
}

// AsahiItem is the fundamental Asahi data type containing a headline, link, date, etc.
type AsahiItem struct {
	Title   string `xml:"title"`
	TitleEN string
	Link    string `xml:"link"`
	Date    string `xml:"date"`
}

// SetFeed should be used where Asahi or NHK is passed with it's appropriate URL. This will then
// set those structs with their appropriate values.
func SetFeed(feed interface{}, url string) error {
	// First fetch the RSS feed.
	resp, err := fetchRSS(url)
	//resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		return err
	}

	// Turn the xml response into a struct.
	err = xml.Unmarshal(resp, feed)
	if err != nil {
		return err
	}

	// Now turn the struct's items into a json []byte.
	// article, err := json.Marshal(feed.XMLCh.Items)
	// if err != nil {
	// 	return "", err
	// }
	return nil
}

// fetchFeed get's the response body of from the passed url and returns a []byte of that
// resposne. Basically just a fancy wrapper for http.Get.
func fetchRSS(url string) ([]byte, error) {
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

func TranslateTitle(n []NHKItem, a []AsahiItem) {
	for i, item := range n {
		if i == 3 {
			break
		}
		var err error
		n[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
		if err != nil {
			fmt.Println(err)
		}
	}

	for i, item := range a {
		if i == 3 {
			break
		}
		var err error
		a[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
		if err != nil {
			fmt.Println(err)
		}
	}
}
