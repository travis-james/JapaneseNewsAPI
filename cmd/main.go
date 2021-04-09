package main

import (
	"fmt"

	"github.com/travis-james/RSSReader/readfeeds"
)

// For NHK.
func main() {
	a := readfeeds.GetNHK()
	fmt.Println(a)
	// b := readfeeds.GetAsahi()
	// fmt.Println(b)
}
