package main

import (
	"fmt"
	"log"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mytranslate"
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
	a := News{}
	var err error
	a.Asahi, err = readfeeds.GetAsahi()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	for i, item := range a {
		a[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
		if err != nil {
			fmt.Println(err)
		}
	}
	for _, item := range a {
		fmt.Println(item)
	}
	// b, err := readfeeds.GetAsahi()
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// fmt.Println(b)
	// c := "今日は、ギターを買うつもりだと思う。"
	// d, err := mytranslate.TranslateJP(c)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(d)
	c, err := mytwitter.GetTrends()
	if err != nil {
		log.Fatalln(err)
	}

	for _, trend := range c.Trends {
		fmt.Println(trend)
	}

}
