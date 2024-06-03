package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func DataBase() *sqlx.DB {
	return db
}

func Connection() {
	var err error
	dns := os.Getenv("DNS")

	db, err = sqlx.Connect("postgres", dns)
	if err != nil {
		log.Fatalln(err)
	}

}
