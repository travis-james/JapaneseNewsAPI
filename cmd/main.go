package main

import (
	"fmt"

	"github.com/travis-james/RSSReader/readfeeds"
)

// For NHK.
func main() {
	// a := readfeeds.GetNHK()
	// for _, item := range a.XMLCh.Items {
	// 	fmt.Println(item.Title, item.Link, item.Date)
	// }
	b := readfeeds.GetAsahi()
	for _, item := range b.Items {
		fmt.Println(item.Title, item.Link, item.Date)
	}
}
