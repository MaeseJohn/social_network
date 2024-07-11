package infrastructure

import (
	jwtservice "social_media/utilities/jwt"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	e.POST("/users", CreatUserHandler(pgRepository))
	e.GET("/users", FindAllUsersHandler(pgRepository))
	jwtService := jwtservice.NewJWTService("secretkey")
	e.POST("/login", LoginUserHandler(jwtService, pgRepository))
}
