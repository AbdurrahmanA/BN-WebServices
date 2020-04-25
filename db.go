package main

import (
	"log"

	"github.com/go-bongo/bongo"
)

var connection = conDb()

//var routers = router()

func conDb() *bongo.Connection {
	config := &bongo.Config{
		ConnectionString: "localhost",
		Database:         "BN",
	}
	connection, err := bongo.Connect(config)
	if err != nil {
		log.Fatal(err)
	}
	return connection
}

/*
func router() *http.ServeMux {
	mux := http.NewServeMux()
	return mux
}
func handler() http.Handler {
	handler := cors.Default().Handler(router())
	return handler
}
*/
