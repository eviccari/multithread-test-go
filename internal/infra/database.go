package infra

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/eviccari/multithread-test-go/configs"
)

func GetDB() (db *sql.DB, err error) {
	log.Printf("DATABASE_URI: %s", getURI())
	db, err = sql.Open(configs.DBEngine, getURI())
	if err != nil {
		log.Fatalf("error on connect to database: %s", err.Error())
	}

	db.SetMaxOpenConns(500)
	db.SetMaxIdleConns(500)

	return
}

func getURI() (URI string) {
	URI = fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		configs.DBUser,
		configs.DBPassword,
		configs.DBHostName,
		configs.DBPort,
		configs.DBName,
	)

	return
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Fatalf("error on close db: %s", err.Error())
	}
}
