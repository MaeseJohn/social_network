package infrastructure

import (
	"net/http"
	"social_media/user/domain"

	"github.com/labstack/echo/v4"
)

var m = map[error]int{
	domain.ErrInternalServerError: http.StatusInternalServerError,
	domain.ErrInvalidCredentials:  http.StatusUnauthorized,
	domain.ErrNotFound:            http.StatusNotFound,
	domain.ErrUnprocessableEntity: http.StatusUnprocessableEntity,
	domain.ErrBadRequest:          http.StatusBadRequest,
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	c.JSON(m[err], map[string]interface{}{
		"code":    m[err],
		"message": err.Error(),
	})
}
