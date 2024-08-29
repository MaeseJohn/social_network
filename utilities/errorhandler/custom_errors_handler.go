package customerros

import (
	"maps"
	follow "social_media/follow/infrastructure"
	user "social_media/user/infrastructure"

	"github.com/labstack/echo/v4"
)

var m = map[error]int{}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	maps.Copy(m, user.ErrorsMap)
	maps.Copy(m, follow.ErrorsMap)
	c.JSON(m[err], map[string]interface{}{
		"code":    m[err],
		"message": err.Error(),
	})
}
