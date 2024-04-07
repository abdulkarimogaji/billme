package main

import (
	"log"

	"github.com/abdulkarimogaji/billme/config"
	"github.com/abdulkarimogaji/billme/db"
	"github.com/abdulkarimogaji/billme/server"
)

func main() {
	// load config
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %s", err)
	}

	// connect db
	err = db.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to db %s", err)
	}

	// start server
	err = server.StartServer()
	if err != nil {
		log.Fatalf("Failed to start server %s", err)
	}
}
