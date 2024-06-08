package infrastructure

import (
	"social_media/db"
	userdomain "social_media/user/domain"
)

type PostgresRepository struct {
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
