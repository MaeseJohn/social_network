package infrastructure

import (
	"net/http"
	"social_media/follow/application"
	"social_media/follow/domain"

	"github.com/labstack/echo/v4"
)

type params struct {
	FollowedId string
}

func FollowHandler(rep domain.FollowRepository, jwtService application.JWTService) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		var params params
		err := ctx.Bind(&params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		reqToken := ctx.Request().Header.Get("Authorization")
		err = application.FollowUC(reqToken, params.FollowedId, rep, jwtService)
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
