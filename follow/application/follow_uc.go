package application

import (
	"social_media/follow/domain"

	"github.com/golang-jwt/jwt/v5"
)

func FollowUC(receiverId string, rep domain.FollowRepository, claims jwt.MapClaims) (error, string) {

	senderId := claims["user_id"].(string)
	exist, err := rep.CheckFollowExists(senderId, receiverId)
	if err != nil {
		return err, ""
	}
	if exist {
		return domain.ErrFollowExist, ""
	}

	private, err := rep.GetUserPrivacity(receiverId)
	if err != nil {
		return err, err.Error()
	}

	if private {
		err := FollowRequestUC(receiverId, senderId, rep)
		if err != nil {
			return err, ""
		}
		return nil, "Follow request succesfully"
	}

	follow := domain.NewFollow(senderId, receiverId)
	err = rep.SaveFollow(follow)
	if err != nil {
		return err, err.Error()
	}

	return nil, "Follow succesfully"
}
