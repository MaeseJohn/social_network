package infrastructure

import (
	"net/http"
	"social_media/user/domain"
)

var ErrorsMap = map[error]int{
	domain.ErrInternalServerError: http.StatusInternalServerError,
	domain.ErrInvalidCredentials:  http.StatusUnauthorized,
	domain.ErrNotFound:            http.StatusNotFound,
	domain.ErrUnprocessableEntity: http.StatusUnprocessableEntity,
	domain.ErrBadRequest:          http.StatusBadRequest,
}
