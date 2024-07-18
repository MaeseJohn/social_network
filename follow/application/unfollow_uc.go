package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func UnfollowUC(followedId string, rep domain.FollowRepository, claims jwt.MapClaims) error {
	followerId := claims["user_id"].(string)
	err := rep.DeleteFollow(followerId, followedId)
	if err != nil {
		return err
	}

	return nil
}
