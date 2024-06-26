package validator

import (
	"social_media/user/domain"

	"github.com/go-playground/validator/v10"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.Validator.Struct(i)
	if err != nil {
		return domain.ErrBadRequest
	}
	return nil
}
