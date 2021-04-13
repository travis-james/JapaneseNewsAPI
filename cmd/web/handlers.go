package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	info := []byte(`Welcome to JapaneseNewsAPI.\n
	updatenews => update the database with the latest news.\n
	/getnews => get news for that given date.`)
	w.Write(info)
}

func (app *application) updatenews(w http.ResponseWriter, r *http.Request) {
	// Get NHK.
	n := &readfeeds.NHK{}
	err := readfeeds.SetFeed(n, nhkURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range n.XMLCh.Items {
	// 	fmt.Println(item)
	// }

	// Get Asahi.
	a := &readfeeds.Asahi{}
	err = readfeeds.SetFeed(a, asahiURL)
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	// for _, item := range a.Items {
	// 	fmt.Println(item)
	// }

	// Translate Asahi & NHK.
	readfeeds.TranslateTitle(n.XMLCh.Items, a.Items)
	// for i, item := range a.Items {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(item)
	// }
	// for i, item := range n.XMLCh.Items {
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(item)
	// }

	// Get twitter trends.
	c, err := mytwitter.GetTrends()
	if err != nil {
		log.Fatalln(err)
	}
	// for _, trend := range c.Trends {
	// 	fmt.Println(trend)
	// }

	todaysnews := News{
		NHK:   n.XMLCh.Items,
		Asahi: a.Items,
		Twit:  c,
		Date:  time.Now(),
		ID:    2,
	}
	//fmt.Println(todaysnews)

	// JSON?
	// jnews, err := json.Marshal(todaysnews)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(jnews))

	// For DB.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("jpnews").Collection("day")
	//ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, todaysnews)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}

func (app *application) getnews(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from getnews"))
}
