package main

import (
	"fmt"
	"log"

	"github.com/travis-james/JapaneseNewsAPI/mytranslate"
	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
)

// For NHK.
func main() {
	a, err := readfeeds.GetAsahi()
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
}
