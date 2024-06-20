package infrastructure

import (
	"log"
	"net/http"
	userapp "social_media/user/application"
	"social_media/user/domain"
	userdomain "social_media/user/domain"

	"github.com/labstack/echo/v4"
)

// Get data, validate and call the UC (separation of concerns)
func CreatUserHandler(r userdomain.UserRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var p userapp.CreateUserParams

		if err := ctx.Bind(&p); err != nil {
			log.Println(err)
			return domain.ErrUnprocessableEntity
		}

		if err := ctx.Validate(p); err != nil {
			log.Println(err)
			return domain.ErrUnprocessableEntity
		}

		if err := userapp.CreateUserUC(&p, r); err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusCreated,
			"message": "User created",
		})
	}
}
