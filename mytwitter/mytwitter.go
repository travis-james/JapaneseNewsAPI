package mytwitter

import (
	"context"
	"fmt"
	"os"
	"time"

	translate "cloud.google.com/go/translate/apiv3"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
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
	//wg             = &sync.WaitGroup{}
)

func main() {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	twitcliA := twitter.NewClient(httpClient)
	trends()
	//twitcliB := twitter.NewClient(httpClient)

	// Create a google translate client.
	// ctx := context.Background()
	// tc, err := translate.NewTranslationClient(ctx, option.WithCredentialsFile("translation-api-project-307416-8cf7f95bb9a6.json"))
	// if err != nil {
	// 	fmt.Println("Client failed.")
	// 	panic(err)
	// }

	//wg.Add(2)
	//go trends(ctx, twitcliA, tc)
	//go retweet(twitcliB, "golang")
	//wg.Wait()
}

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("Missing environment variable " + name)
	}
	return v
}

func retweet(client *twitter.Client, hashtag string) {
	//defer wg.Done()

	// Crack open a stream of tweets.
	stream, err := client.Streams.Filter(&twitter.StreamFilterParams{
		Language: []string{"ja"},
		Track:    []string{hashtag},
	})
	if err != nil {
		panic(err)
	}
	defer stream.Stop()

	for val := range stream.Messages {
		// Because I don't know what could be coming from the chan interface{}
		// assert that the value is a tweet.
		t, ok := val.(*twitter.Tweet)
		if !ok {
			fmt.Println("bad val")
			continue
		}

		// For some reason t.Retweeted is always false, so my understanding of it
		// is wrong, but I can't find another parameter where I can check if I've
		// retweeted, so I'm going by name, I guess.
		if t.User.ScreenName == me || t.Retweeted {
			continue
		}

		_, _, err := client.Statuses.Retweet(t.ID, nil)
		if err != nil {
			fmt.Printf("error in retweeting: %d\n %v\n", t.ID, err)
			continue
		}
		fmt.Printf("retweeted: %d\n", t.ID)
	}
}

func trends(client *twitter.Client) {

	for {
		// Get current time/date in PST
		loc, _ := time.LoadLocation("America/Los_Angeles")
		now := time.Now().In(loc)
		text := fmt.Sprintf("The current time is: %s\nCurrent trends in Japan are:\n", now)

		// Get a trend list of the latest trends.
		tl, _, err := client.Trends.Place(jpWOEID, nil)
		if err != nil {
			panic(err)
		}

		// Get 5 trends and append them to text.
		for i := 0; i < 5; i++ {
			tt, err := translateJP(ctx, tl[0].Trends[i].Name, tc)
			if err != nil {
				fmt.Printf("Error in translate:\n %v\n", err)
				continue
			}
			text += "‣ " + tt + "\n"
		}
		fmt.Println("the text: " + text)
		_, _, err = client.Statuses.Update(text, nil)
		if err != nil {
			fmt.Printf("Error in tweeting trends:\n %v\n", err)
		}
		time.Sleep(12 * time.Hour)
	}
}

func translateJP(ctx context.Context, text string, tc *translate.TranslationClient) (string, error) {
	req := &translatepb.TranslateTextRequest{
		Contents:           []string{text},
		SourceLanguageCode: "ja",
		TargetLanguageCode: "en",
		Parent:             "projects/translation-api-project-307416",
	}
	resp, err := tc.TranslateText(ctx, req)
	if err != nil {
		fmt.Println("Translate failed.")
		return "", err
	}
	return resp.Translations[0].TranslatedText, nil
}
