package infrastructure

import (
	"net/http"
	"social_media/follow/domain"
)

var ErrorsMap = map[error]int{
	domain.ErrAlredyExist:         http.StatusBadRequest,          //400
	domain.ErrUnauthorized:        http.StatusUnauthorized,        //401
	domain.ErrNotFound:            http.StatusNotFound,            //404
	domain.ErrUnprocessableEntity: http.StatusUnprocessableEntity, //422
	domain.ErrInternalServerError: http.StatusInternalServerError, //500
}
