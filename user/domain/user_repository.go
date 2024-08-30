package domain

type UserRepository interface {
	Save(u *User) error
	GetUsers() ([]User, error)
	GetUser(email string) (*User, error)
}
