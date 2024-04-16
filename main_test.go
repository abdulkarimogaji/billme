package main_test

import (
	"log"
	"testing"

	"github.com/abdulkarimogaji/billme/config"
	"github.com/abdulkarimogaji/billme/db"
)

func TestMain(t *testing.M) {
	// load config
	err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config %s", err)
	}

	// connect db
	err = db.ConnectTestDB()
	if err != nil {
		log.Fatalf("Failed to connect to db %s", err)
	}

}
