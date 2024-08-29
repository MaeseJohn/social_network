package domain

type UserRepository interface {
	Save(u *User) error
	GetUsers() ([]string, error)
	GetUser(email string) (*User, error)
}
