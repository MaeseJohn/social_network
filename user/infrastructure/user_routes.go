package infrastructure

import (
	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	e.POST("/users", CreatUserHandler(pgRepository))
	e.GET("/users", FindAllUsersHandler(pgRepository))
	jwtService := NewJWTService("secretkey")
	e.POST("/login", LoginUserHandler(jwtService, pgRepository))
}
