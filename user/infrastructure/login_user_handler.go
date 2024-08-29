package infrastructure

import (
	"net/http"
	"social_media/user/application"
	"social_media/user/domain"

	"github.com/labstack/echo/v4"
)

type LoginParamas struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
}

func LoginUserHandler(jwtService application.JWTService, rep domain.UserRepository) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var params LoginParamas
		err := ctx.Bind(&params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		err = ctx.Validate(params)
		if err != nil {
			return domain.ErrUnprocessableEntity
		}

		token, err := application.LoginUC(params.Email, params.Password, rep, jwtService)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, map[string]interface{}{
			"code":    http.StatusOK,
			"message": "Login successfully",
			"token":   token,
		})
	}
}
