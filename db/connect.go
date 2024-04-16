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
var TestStorage DBStorage

func ConnectDB() error {
	dbConfig := mysql.Config{
		User:                 config.AppConfig.DB_USER,
		Passwd:               config.AppConfig.DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 config.AppConfig.DB_ADDR,
		DBName:               "billme",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	log.Println(dbConfig.FormatDSN())
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return err
	}
	Storage.DB = db
	return db.Ping()
}

func ConnectTestDB() error {
	dbConfig := mysql.Config{
		User:                 config.AppConfig.DB_USER,
		Passwd:               config.AppConfig.DB_PASSWORD,
		Net:                  "tcp",
		Addr:                 config.AppConfig.DB_ADDR,
		DBName:               "billme_test",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
	log.Println(dbConfig.FormatDSN())
	// Get a database handle.
	var err error
	db, err := sql.Open("mysql", dbConfig.FormatDSN())
	if err != nil {
		return err
	}
	TestStorage.DB = db
	return db.Ping()
}

func (s *DBStorage) Ping() error {
	return s.DB.Ping()
}
