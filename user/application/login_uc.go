package application

import (
	"social_media/user/domain"
)

type JWTService interface {
	CreateToken(user *domain.User) (string, error)
}

func LoginUC(email, password string, rep domain.UserRepository, jwt JWTService) (string, error) {
	user, err := rep.GetUser(email)
	if err != nil {
		return "", err
	}

	if !user.ValidatePassword(password) {
		return "", domain.ErrInvalidCredentials
	}

	token, err := jwt.CreateToken(user)
	if err != nil {
		return "", err //Check this error
	}

	return token, nil
}
