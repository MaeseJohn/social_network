package infrastructure

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterFollowRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	f := e.Group("/follow")
	f.Use(echojwt.JWT([]byte("secretkey")))
	f.POST("/follow", FollowHandler(pgRepository))
}
