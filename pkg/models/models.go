package models

import (
	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
)

// News is the struct containing all the translated news, twitter trends,
// and date.
type News struct {
	NHK   []mynews.NHKItem
	Asahi []mynews.AsahiItem
	Twit  []mytwitter.TTrend
	Date  string
}
