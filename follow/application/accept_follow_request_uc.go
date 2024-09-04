package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func AcceptFollowRequestUC(rep domain.FollowRepository, claims jwt.MapClaims, senderId string) error {
	receiverId := claims["user_id"].(string)
	err := FollowResponseValidations(rep, senderId, receiverId)
	if err != nil {
		return err
	}

	follow := domain.NewFollow(senderId, receiverId)
	rep.AcceptFollowRequest(follow, senderId, receiverId)
	return nil
}
