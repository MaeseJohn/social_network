package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func FollowRequestUC(receiverId string, rep domain.FollowRepository, claims jwt.MapClaims) error {
	senderId := claims["user_id"].(string)
	exist, err := rep.CheckFriendRequestExists(senderId, receiverId)
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrAlredyExist
	}

	followRequest := domain.NewFollowRequest(senderId, receiverId)
	err = rep.RequestFollow(followRequest)
	if err != nil {
		return err
	}

	return nil
}
