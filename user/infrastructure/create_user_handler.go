package infrastructure

import (
	"fmt"
	"net/http"
	"social_media/user/application"
	"social_media/user/domain"

	"github.com/labstack/echo/v4"
)

// Get data, validate and call the UC (separation of concerns)
func CreatUserHandler(r domain.UserRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var p application.CreateUserParams

		err := ctx.Bind(&p)
		if err != nil {
			fmt.Println(err)
			return domain.ErrUnprocessableEntity
		}

		err = ctx.Validate(p)
		if err != nil {
			fmt.Println(p)
			fmt.Println("validate")
			fmt.Println(err)
			return domain.ErrUnprocessableEntity
		}

		err = application.CreateUserUC(&p, r)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusCreated,
			"message": "User created",
		})
	}
}
