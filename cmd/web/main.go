package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/travis-james/JapaneseNewsAPI/pkg/mymongo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// application struct allows dependency injection.
type application struct {
	news *mymongo.NewsModel
	ctx  context.Context
}

var (
	dbUser = getenv("DB_USER")
	dbPwd  = getenv("DB_PASS")
	dbName = getenv("DB_AT")
)

func getenv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		panic("Missing environment variable " + name)
	}
	return v
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	dsn := "mongodb+srv://" + dbUser + ":" + dbPwd + "@" + dbName

	// For DB.
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(dsn)
	dbc, err := mongo.Connect(c, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		news: &mymongo.NewsModel{DB: dbc},
		ctx:  c,
	}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %v\n", *addr)
	err = srv.ListenAndServe()
	log.Fatalln(err)
}
