package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string
	Name     string
	LastName string
	Email    string
	Password string
	Age      string
}

func (u *User) HashSaltPassword() error {
	password := []byte(u.Password)
	//Hashing the password with the default cost of 10
	//The salt is automatically (and randomly) generated upon hashing a password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	return err
}

// Returns true if the password is correct, false if not
func (u *User) ValidatePassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewUser(id, name, lastName, email, password, age string) (*User, error) {
	user := User{
		Id:       id,
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: password,
		Age:      age,
	}

	if err := user.HashSaltPassword(); err != nil {
		return nil, err
	}
	return &user, nil
}
