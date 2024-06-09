package application

import (
	"fmt"
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

	fmt.Println("getuser works")

	if !user.ValidatePassword(password) {
		return "", fmt.Errorf("Invalid credetials")
	}

	fmt.Println("validatepassword works")

	token, err := jwt.CreateToken(user)
	if err != nil {
		return "", err
	}

	return token, nil
}
