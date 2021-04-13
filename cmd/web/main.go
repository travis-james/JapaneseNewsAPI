package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

// application struct allows dependency injection. For example,
// by using this, my handlers could have access to using
//  logs I declare in main. As of now it's just for DB.
type application struct {
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	app := &application{}

	srv := &http.Server{
		Addr:    *addr,
		Handler: app.routes(),
	}
	fmt.Printf("Listening on port %v\n", *addr)
	err := srv.ListenAndServe()
	log.Fatalln(err)
}
