package mongo

import (
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

var id = 0

func (mc *mongo.Client) Insert(n []mynews.NHKItem, a []mynews.AsahiItem, t []mytwitter.TTrend) (int, error) {
	date := time.Now().Format("2006-01-02")
	todaysnews := models.News{
		NHK:   n.XMLCh.Items,
		Asahi: a.Items,
		Twit:  t,
		Date:  date,
		ID:    id,
	}
	id++
}
