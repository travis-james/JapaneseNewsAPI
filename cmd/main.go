package main

import (
	"fmt"
	"log"

	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
)

// For NHK.
func main() {
	// a, err := readfeeds.GetAsahi()
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// for i, item := range a {
	// 	a[i].TitleEN, err = mytranslate.TranslateJP(item.Title)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// for _, item := range a {
	// 	fmt.Println(item)
	// }
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
	a, err := mytwitter.GetTrends()
	if err != nil {
		log.Fatalln(err)
	}

	for _, trend := range a.Trends {
		fmt.Println(trend)
	}

}
