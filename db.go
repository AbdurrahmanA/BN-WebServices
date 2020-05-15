package main

import (
	"sync"

	"github.com/go-bongo/bongo"
)

var connection *bongo.Connection
var lock = &sync.Mutex{}

func conDb() *bongo.Connection {
	if connection == nil {
		lock.Lock()
		if connection == nil {
			config := &bongo.Config{
				ConnectionString: "localhost",
				Database:         "BN",
			}
			connection, err := bongo.Connect(config)
			if err != nil {
				connection = nil
				return connection
			}
			return connection
		}
		defer lock.Unlock()
		return connection
	}
	return connection

}
