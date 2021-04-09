package mytwitter

import (
	"fmt"
	"os"
	"time"

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

type TTrends struct {
	Trends []TTrend
	Time   time.Time
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

func GetTrends() (*TTrends, error) {
	// Get Credentials for twitter.
	client := getCredentials()
	// Get current time/date in PST
	loc, _ := time.LoadLocation("America/Los_Angeles")
	now := time.Now().In(loc)
	retval := &TTrends{
		Trends: make([]TTrend, 5),
		Time:   now,
	}

	// Get a trend list of the latest trends.
	tl, _, err := client.Trends.Place(jpWOEID, nil)
	if err != nil {
		return nil, err
	}

	// Get 5 trends and append them to text.
	for i := 0; i < 5; i++ {
		retval.Trends[i].Trend = tl[0].Trends[i].Name
		retval.Trends[i].TrendEN, err = mytranslate.TranslateJP(tl[0].Trends[i].Name)
		if err != nil {
			fmt.Println(err)
		}
	}
	return retval, nil
}
