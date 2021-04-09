package main

import (
	"fmt"
	"os"
	"time"

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
)

func main() {
	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	// http.Client will automatically authorize Requests
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	twitcliA := twitter.NewClient(httpClient)
	trends(twitcliA)
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

func trends(client *twitter.Client) {
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
	for i := 0; i < 10; i++ {
		text += "‣ " + tl[0].Trends[i].Name + "\n"
	}
	fmt.Println("the text: " + text)
}

// func translateJP(ctx context.Context, text string, tc *translate.TranslationClient) (string, error) {
// 	req := &translatepb.TranslateTextRequest{
// 		Contents:           []string{text},
// 		SourceLanguageCode: "ja",
// 		TargetLanguageCode: "en",
// 		Parent:             "projects/translation-api-project-307416",
// 	}
// 	resp, err := tc.TranslateText(ctx, req)
// 	if err != nil {
// 		fmt.Println("Translate failed.")
// 		return "", err
// 	}
// 	return resp.Translations[0].TranslatedText, nil
// }
