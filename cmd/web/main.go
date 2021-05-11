package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/pkg/mymongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// application struct allows dependency injection. For example,
// by using this, my handlers could have access to using
//  logs I declare in main. As of now it's just for DB.
type application struct {
	news *mymongo.NewsModel
	ctx  context.Context
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	dsn := flag.String("dsn", "mongodb://localhost:27017", "The data source name for the database")
	flag.Parse()

	// For DB.
	c, _ := context.WithTimeout(context.Background(), 10*time.Second)
	clientOptions := options.Client().ApplyURI(*dsn)
	dbc, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		client: dbc,
		ctx:    c,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %v\n", *addr)
	err = srv.ListenAndServe()
	log.Fatalln(err)
}
