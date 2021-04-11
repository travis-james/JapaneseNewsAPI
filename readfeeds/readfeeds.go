package readfeeds

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytranslate"
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
	Title   string `xml:"title"`
	TitleEN string
	Link    string `xml:"link"`
	Date    string `xml:"pubDate"`
}

// GetAsahi goes to the Asahi RSS feed and turns that xml response into a struct. The part of the struct containing
// a slice of headlines/links/dats is returned.
func GetNHK() ([]NHKItem, error) {
	// First fetch the RSS feed.
	resp, err := fetchFeed(nhkURL)
	//resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		return nil, err
	}

	// Turn the xml response into a struct.
	feed := &NHK{}
	err = xml.Unmarshal(resp, feed)
	if err != nil {
		return nil, err
	}

	// Translate the headlines.
	feed.XMLCh.translatetitle()

	// Now turn the struct's items into a json []byte.
	// article, err := json.Marshal(feed.XMLCh.Items)
	// if err != nil {
	// 	return "", err
	// }
	return feed.XMLCh.Items, nil
}

type Asahi struct {
	XMLName xml.Name    `xml:"RDF"`
	Items   []AsahiItem `xml:"item"`
}

type AsahiItem struct {
	Title   string `xml:"title"`
	TitleEN string
	Link    string `xml:"link"`
	Date    string `xml:"date"`
}

// GetAsahi goes to the Asahi RSS feed and turns that xml response into a struct. The part of the struct containing
// a slice of headlines/links/dats is returned.
func GetAsahi() ([]AsahiItem, error) {
	// First fetch the RSS feed.
	//resp, err := fetchFeed(asahiURL)
	resp, err := ioutil.ReadFile("ex.rdf")
	if err != nil {
		return nil, err
	}

	// Turn the xml response into a struct.
	feed := &Asahi{}
	err = xml.Unmarshal(resp, feed)
	if err != nil {
		return nil, err
	}

	// Translate the headlines.
	feed.translatetitle()

	// Now turn the struct's items into a json []byte.
	// article, err := json.Marshal(feed.Items)
	// if err != nil {
	// 	return "", err
	// }
	return feed.Items, nil
}

// fetchFeed get's the response body of from the passed url and returns a []byte of that
// resposne. Basically just a fancy wrapper for http.Get.
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

type Headline interface {
	translatetitle()
}

func (n *NHKChannel) translatetitle() {
	for i, item := range n.Items {
		// if i == 3 {
		// 	return
		// }
		var err error
		n.Items[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (a *Asahi) translatetitle() {
	for i, item := range a.Items {
		// if i == 3 {
		// 	return
		// }
		var err error
		a.Items[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
		if err != nil {
			fmt.Println(err)
		}
	}
}
