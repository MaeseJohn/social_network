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
	rep.AcceptFollowRequest(follow, senderId, receiverId)
	return nil
}
