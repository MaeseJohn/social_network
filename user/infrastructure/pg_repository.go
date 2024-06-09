package infrastructure

import (
	"fmt"
	"social_media/db"
	userdomain "social_media/user/domain"
)

type PostgresRepository struct {
}

type DBUser struct {
	Id       string
	Name     string
	LastName string `db:"last_name"`
	Email    string
	Password string
	Age      string
}

func (user *DBUser) dbUserToDomainUser() *userdomain.User {
	return &userdomain.User{
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

func (*PostgresRepository) Save(u *userdomain.User) error {
	_, err := db.DataBase().
		NamedExec("INSERT INTO users (id, name, last_name, email, password, age) VALUES (:id, :name, :lastname, :email, :password, :age)", u)
	if err != nil {
		return err
	}

	return nil
}

func (*PostgresRepository) GetUsers() ([]string, error) {
	var users []string
	err := db.DataBase().
		Select(&users, "SELECT name FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (*PostgresRepository) GetUser(email string) (*userdomain.User, error) {
	var user DBUser
	err := db.DataBase().
		Get(&user, "SELECT * FROM users WHERE email=$1", email)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user.dbUserToDomainUser(), nil
}
