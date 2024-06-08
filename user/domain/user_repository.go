package domain

type UserRepository interface {
	Save(u *User) error
}
