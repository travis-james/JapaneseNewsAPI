package models

import (
	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
)

// News is the struct containing all the translated news, twitter trends,
// and date. An id is included for my own dev/test purposes.
type News struct {
	NHK   []mynews.NHKItem
	Asahi []mynews.AsahiItem
	Twit  []mytwitter.TTrend
	Date  string
	ID    int
}
