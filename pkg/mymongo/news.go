package mymongo_test

import (
	"context"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/mynews"
	"github.com/travis-james/JapaneseNewsAPI/mytwitter"
	"github.com/travis-james/JapaneseNewsAPI/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewsModel is just a type to wrap the mongo.Client, so I can add more methods.
type NewsModel struct {
	DB *mongo.Client
}

var id = 0

func (nm *NewsModel) Insert(n mynews.NHK, a mynews.Asahi, t []mytwitter.TTrend) (*mongo.InsertOneResult, error) {
	date := time.Now().Format("2006-01-02")
	todaysnews := models.News{
		NHK:   n.XMLCh.Items,
		Asahi: a.Items,
		Twit:  t,
		Date:  date,
		ID:    id,
	}

	collection := nm.DB.Database("jpnews").Collection("day")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, err := collection.InsertOne(ctx, todaysnews)
	if err != nil {
		return nil, err
	}
	return result, nil
}
