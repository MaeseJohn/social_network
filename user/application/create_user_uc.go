package application

import (
	"social_media/user/domain"

	"github.com/google/uuid"
)

type CreateUserParams struct {
	Name     string `validate:"required,alpha"`
	LastName string `validate:"required,alpha"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
	Age      string `validate:"required,datetime=2006-01-02"`
	Private  bool
}

func CreateUserUC(u *CreateUserParams, rep domain.UserRepository) error {

	user, err := domain.NewUser(uuid.NewString(), u.Name, u.LastName, u.Email, u.Password, u.Age, u.Private)
	if err != nil {
		return err
	}

	err = rep.Save(user)
	if err != nil {
		return err
	}

	return nil
}
