package main

import (
	"social_media/db"
	"social_media/db/migrations"
	followinfrastructure "social_media/follow/infrastructure"
	userinfrastructure "social_media/user/infrastructure"
	customerros "social_media/utilities/errorhandler"
	"social_media/utilities/validator"

	goPlayground "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
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

	e := echo.New()
	e.Validator = &validator.CustomValidator{Validator: goPlayground.New()}
	e.HTTPErrorHandler = customerros.CustomHTTPErrorHandler
	userinfrastructure.RegisterUserRoutes(e)
	followinfrastructure.RegisterFollowRoutes(e)
	e.Logger.Fatal(e.Start(":3000"))
}
