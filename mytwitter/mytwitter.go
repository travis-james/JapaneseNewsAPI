package mytwitter

import (
	"fmt"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/travis-james/JapaneseNewsAPI/mytranslate"
)

const (
	jpWOEID = 23424856
	me      = "JPtoEN_Bot"
)

var (
	consumerKey    = getenv("TWITTER_CONSUMER_KEY")
	consumerSecret = getenv("TWITTER_CONSUMER_SECRET")
	accessToken    = getenv("TWITTER_ACCESS_TOKEN")
	accessSecret   = getenv("TWITTER_ACCESS_TOKEN_SECRET")
)

type TTrend struct {
	Trend   string
	TrendEN string
}

func getCredentials() *twitter.Client {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	twitcli := twitter.NewClient(httpClient)
	return twitcli
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("Missing environment variable " + name)
	}
	return v
}

// Why return pointer instead of value? So I can return nil....
func GetTrends() ([]TTrend, error) {
	// Get Credentials for twitter.
	client := getCredentials()
	// Get current time/date in PST
	//loc, _ := time.LoadLocation("America/Los_Angeles")
	retval := make([]TTrend, 5)

	// Get a trend list of the latest trends.
	tl, _, err := client.Trends.Place(jpWOEID, nil)
	if err != nil {
		return nil, err
	}

	// Get 5 trends, translate, and append them to the slice.
	for i := 0; i < 5; i++ {
		retval[i].Trend = tl[0].Trends[i].Name
		retval[i].TrendEN, err = mytranslate.TranslateJP(tl[0].Trends[i].Name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return retval, nil
}
