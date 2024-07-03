package infrastructure

import (
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
