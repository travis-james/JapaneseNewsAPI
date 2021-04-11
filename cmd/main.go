package main

import (
	"fmt"
	"log"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
)

type News struct {
	NHK   []readfeeds.NHKItem
	Asahi []readfeeds.AsahiItem
	Twit  *mytwitter.TTrends
	Date  time.Time
}

// For NHK.
func main() {
	asahinews, err := readfeeds.GetNHK()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	for _, item := range asahinews {
		fmt.Println(item)
	}
	// c, err := mytwitter.GetTrends()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// for _, trend := range c.Trends {
	// 	fmt.Println(trend)
	// }

}
