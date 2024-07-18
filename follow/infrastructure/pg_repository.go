package infrastructure

import (
	"fmt"
	"social_media/db"
	"social_media/follow/domain"
)

type PostgresRepository struct {
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}

func (*PostgresRepository) SaveFollow(follow *domain.Follow) error {
	_, err := db.DataBase().
		NamedExec("INSERT INTO follows (follower_id, followed_id, follow_date) VALUES (:followerid, :followedid, :followdate)", follow)

	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (*PostgresRepository) DeleteFollow(followerid, followedid string) error {

	_, err := db.DataBase().
		Exec("DELETE FROM follows WHERE follower_id=$1 AND followed_id=$2", followerid, followedid)

	if err != nil {
		fmt.Println(err)
		return domain.ErrInternalServerError
	}

	return nil
}
