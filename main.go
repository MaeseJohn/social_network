package main

import (
	"social_media/db"
	"social_media/db/migrations"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Fail loading .env file")
	}

	//Connect to db and creat de db instance
	db.Connection()

	//Fill data base with migrations
	migrations.SetupDB()

}
