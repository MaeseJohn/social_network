package application

import (
	"social_media/follow/domain"
)

func FollowRequestUC(receiverId, senderId string, rep domain.FollowRepository) error {
	exist, err := rep.CheckFollowRequestExists(senderId, receiverId)
	if err != nil {
		return err
	}

	if exist {
		return domain.ErrRequestExist
	}

	followRequest := domain.NewFollowRequest(senderId, receiverId)
	err = rep.RequestFollow(followRequest)
	if err != nil {
		return err
	}

	return nil
}
