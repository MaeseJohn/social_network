package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func FollowUC(followedId string, rep domain.FollowRepository, claims jwt.MapClaims) error {
	followerId := claims["user_id"].(string)
	follow := domain.NewFollow(followerId, followedId)
	err := rep.SaveFollow(follow)
	if err != nil {
		return err
	}

	return nil
}
