package infrastructure

import (
	"net/http"
	"social_media/follow/application"
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetFollowRequestHandler(rep domain.FollowRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		var followRequests []domain.FollowRequest
		followRequests, err := application.GetFollowRequestUC(rep, claims)
		if err != nil {
			return err
		}
		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Follow requests succesfully found",
			"data":    followRequests,
		})
	}
}
