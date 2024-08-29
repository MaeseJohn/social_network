package infrastructure

import (
	"net/http"
	"social_media/follow/application"
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func FollowHandler(rep domain.FollowRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var params followedid
		err := ctx.Bind(&params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		err = ctx.Validate(params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		private, err := rep.GetUserPrivacity(params.FollowedId)
		if err != nil {
			return err
		}

		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)

		var responseMessage string
		if private {
			err = application.FollowRequestUC(params.FollowedId, rep, claims)
			if err != nil {
				return err
			}
			responseMessage = "Follow request sent"

		} else {
			err = application.FollowUC(params.FollowedId, rep, claims)
			if err != nil {
				return err
			}
			responseMessage = "Follow successfully"
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": responseMessage,
			"data":    params,
		})
	}
}
