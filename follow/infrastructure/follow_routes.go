package infrastructure

import (
	jwtservice "social_media/utilities/jwt"

	"github.com/labstack/echo/v4"
)

func RegisterFollowRoutes(e *echo.Echo) {
	pgRepository := NewPostgresRepository()
	jwtService := jwtservice.NewJWTService("secretkey")
	e.POST("/follow", FollowHandler(pgRepository, jwtService))
}
