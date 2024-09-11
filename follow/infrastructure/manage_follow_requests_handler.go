package infrastructure

import (
	"net/http"
	"social_media/follow/application"
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func ManageFollowRequestsHandler(rep domain.FollowRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		type followRequestResponse struct {
			SenderId string `validate:"required,uuid"`
			Response bool
		}
		var params followRequestResponse
		err := ctx.Bind(&params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		err = ctx.Validate(params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		user := ctx.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		receiverId := claims["user_id"].(string)

		var responseMessage string
		if params.Response {
			err = application.AcceptFollowRequestUC(rep, params.SenderId, receiverId)
			if err != nil {
				return err
			}
			responseMessage = "Follow request acepted"

		} else {
			err = application.DeclineFollowRequestUC(rep, params.SenderId, receiverId)
			if err != nil {
				return err
			}
			responseMessage = "Follow request declined"
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": responseMessage,
			"data":    nil,
		})
	}
}
