package infrastructure

import (
	"net/http"
	userapp "social_media/user/application"
	userdomain "social_media/user/domain"

	"github.com/labstack/echo/v4"
)

// Get data, validate and call the UC (separation of concerns)
func CreatUserHandler(r userdomain.UserRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var p userapp.CreateUserParams

		if err := ctx.Bind(&p); err != nil {
			return ctx.String(http.StatusUnprocessableEntity, err.Error())
		}

		if err := ctx.Validate(p); err != nil {
			return ctx.String(http.StatusUnprocessableEntity, err.Error())
		}

		if err := userapp.CreateUserUC(&p, r); err != nil {
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.String(http.StatusCreated, "Created")
	}
}
