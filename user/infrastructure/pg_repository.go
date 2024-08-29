package infrastructure

import (
	"fmt"
	"social_media/db"
	"social_media/user/domain"
)

type PostgresRepository struct {
}

type DBUser struct {
	Id       string `db:"user_id"`
	Name     string
	LastName string `db:"last_name"`
	Email    string
	Password string
	Age      string
	Private  string
}

func (user *DBUser) dbUserToDomainUser() *domain.User {
	return &domain.User{
		Id:       user.Id,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
		Password: user.Password,
		Age:      user.Age,
	}
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}

func (*PostgresRepository) Save(u *domain.User) error {
	_, err := db.DataBase().
		NamedExec("INSERT INTO users (user_id, name, last_name, email, password, age, private) VALUES (:id, :name, :lastname, :email, :password, :age, :private)", u)
	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (*PostgresRepository) GetUsers() ([]string, error) {
	var users []string
	err := db.DataBase().
		Select(&users, "SELECT name FROM users")
	if err != nil {
		return nil, domain.ErrNotFound
	}

	return users, nil
}

func (*PostgresRepository) GetUser(email string) (*domain.User, error) {
	var user DBUser
	err := db.DataBase().
		Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		fmt.Println(err)
		return nil, domain.ErrNotFound
	}
	return user.dbUserToDomainUser(), nil
}
