package infrastructure

import (
	"net/http"
	"social_media/follow/application"
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type params struct {
	FollowedId string
}

func FollowHandler(rep domain.FollowRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var params params
		err := ctx.Bind(&params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}
		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		err = application.FollowUC(params.FollowedId, rep, claims)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Follow successfully",
			"data":    params,
		})
	}
}
