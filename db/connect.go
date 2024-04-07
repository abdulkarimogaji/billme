package db

import (
	"database/sql"
	"log"

	"github.com/abdulkarimogaji/billme/config"
	"github.com/go-sql-driver/mysql"
)

type DBStorage struct {
	DB *sql.DB
}

var Storage DBStorage

func ConnectDB() error {
	dbConfig := mysql.Config{
		User:                 config.AppConfig.DB_USER,
		Passwd:               config.AppConfig.DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 config.AppConfig.DB_ADDR,
		DBName:               "billme",
		AllowNativePasswords: true,
	}
	log.Println(dbConfig.FormatDSN())
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	Storage.DB = db
	if err != nil {
		return err
	}
	return db.Ping()
}
