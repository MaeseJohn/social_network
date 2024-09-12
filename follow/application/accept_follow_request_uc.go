package application

import (
	"social_media/follow/domain"
)

func AcceptFollowRequestUC(rep domain.FollowRepository, senderId, receiverId string) error {
	err := FollowResponseValidations(rep, senderId, receiverId)
	if err != nil {
		return err
	}

	follow := domain.NewFollow(senderId, receiverId)
	err = rep.AcceptFollowRequest(follow, senderId, receiverId)
	if err != nil {
		return err
	}
	return nil
}
