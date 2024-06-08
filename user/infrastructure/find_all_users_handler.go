package infrastructure

import (
	"net/http"
	userapp "social_media/user/application"
	userdomain "social_media/user/domain"

	"github.com/labstack/echo/v4"
)

func FindAllUsersHandler(rep userdomain.UserRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var users []string
		users, err := userapp.FindAllUsersUC(rep)
		if err != nil {
			return ctx.String(http.StatusNotFound, err.Error())
		}

		return ctx.JSON(http.StatusOK, users)
	}
}
