package infrastructure

import (
	"net/http"
	"social_media/follow/domain"
)

var ErrorsMap = map[error]int{
	domain.ErrUnauthorized:        http.StatusUnauthorized,        //401
	domain.ErrInternalServerError: http.StatusInternalServerError, //500
	domain.ErrUnprocessableEntity: http.StatusUnprocessableEntity, //422
}
