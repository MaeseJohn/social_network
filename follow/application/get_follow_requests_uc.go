package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func GetFollowRequestUC(rep domain.FollowRepository, claims jwt.MapClaims) ([]domain.FollowRequest, error) {
	receiverId := claims["user_id"].(string)
	var followRequests []domain.FollowRequest
	followRequests, err := rep.GetFollowRequests(receiverId)
	if err != nil {
		return nil, err
	}

	return followRequests, nil
}
