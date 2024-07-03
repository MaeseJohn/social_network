package infrastructure

import (
	"github.com/labstack/echo/v4"
)

func RegisterFollowRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	jwtService := NewJWTService("secretkey")
	e.POST("/follow", FollowHandler(pgRepository, jwtService))

}
