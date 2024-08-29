package domain

import "golang.org/x/crypto/bcrypt"

type User struct {
	Id       string
	Name     string
	LastName string
	Email    string
	Password string
	Age      string
	Private  bool
}

func NewUser(id, name, lastName, email, password, age string, private bool) (*User, error) {
	user := User{
		Id:       id,
		Name:     name,
		LastName: lastName,
		Email:    email,
		Password: password,
		Age:      age,
		Private:  private,
	}

	err := user.HashSaltPassword()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) HashSaltPassword() error {
	password := []byte(u.Password)
	//Hashing the password with the default cost of 10
	//The salt is automatically (and randomly) generated upon hashing a password
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	u.Password = string(hashedPassword)
	if err != nil {
		return ErrInternalServerError
	}
	return nil

}

// Returns true if the password is correct, false if not
func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
