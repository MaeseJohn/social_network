package main

import (
	"net/http"
	"social_media/db"
	"social_media/db/migrations"
	"social_media/user/domain"
	userinfrastructure "social_media/user/infrastructure"
	"social_media/validator"

	goPlayground "github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var httpErrors = map[error]int{
	domain.ErrInvalidCredentials: http.StatusUnprocessableEntity,
}

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

	pgRepository := userinfrastructure.NewPostgresRepository()
	e.POST("/users", userinfrastructure.CreatUserHandler(pgRepository))
	e.GET("/users", userinfrastructure.FindAllUsersHandler(pgRepository))
	jwtService := userinfrastructure.NewJWTService("secretkey")
	e.POST("/login", userinfrastructure.LoginUserHandler(jwtService, pgRepository))
	e.Logger.Fatal(e.Start(":3000"))
}
