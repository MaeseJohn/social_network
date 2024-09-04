package application

import (
	"social_media/follow/domain"
)

func FollowResponseValidations(rep domain.FollowRepository, senderId, receiverId string) error {
	exist, err := rep.CheckFollowExists(senderId, receiverId)
	if err != nil {
		return err
	}
	if exist {
		return domain.ErrRequestExist
	}

	exist, err = rep.CheckFollowRequestExists(senderId, receiverId)
	if err != nil {
		return err
	}
	if !exist {
		return domain.ErrNotFound
	}

	return nil
}
