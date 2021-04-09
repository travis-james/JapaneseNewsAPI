package main

import (
	"fmt"
	"log"

	"github.com/travis-james/JapaneseNewsAPI/readfeeds"
)

// For NHK.
func main() {
	a, err := readfeeds.GetNHK()
	if err != nil {
		log.Fatalf("%v\n", err)
	}
	fmt.Println(a)
	// b, err := readfeeds.GetAsahi()
	// if err != nil {
	// 	log.Fatalf("%v\n", err)
	// }
	// fmt.Println(b)
}
