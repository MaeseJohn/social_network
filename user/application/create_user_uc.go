package application

import (
	userdomain "social_media/user/domain"
)

type CreateUserParams struct {
	Id       string `validate:"required,uuid"`
	Name     string `validate:"required,alpha"`
	LastName string `validate:"required,alpha"`
	Email    string `validate:"required,email"`
	Password string `validate:"required,alphanum"`
	Age      string `validate:"required,datetime=2006-01-02"`
}

func CreateUserUC(u *CreateUserParams, rep userdomain.UserRepository) error {

	user, err := userdomain.NewUser(u.Id, u.Name, u.LastName, u.Email, u.Password, u.Age)
	if err != nil {
		return err
	}

	err = rep.Save(user)
	if err != nil {
		return err
	}

	return nil
}
