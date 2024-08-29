package infrastructure

import (
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func RegisterFollowRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	e.POST("/follow", FollowHandler(pgRepository), echojwt.JWT([]byte("secretkey")))
	e.DELETE("/follow", UnfollowHandler(pgRepository), echojwt.JWT([]byte("secretkey")))
	e.GET("/follow", GetFollowRequestHandler(pgRepository), echojwt.JWT([]byte("secretkey")))
}
