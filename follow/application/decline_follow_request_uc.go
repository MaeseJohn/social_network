package application

import (
	"social_media/follow/domain"
)

func DeclineFollowRequestUC(rep domain.FollowRepository, senderId, receiverId string) error {
	err := FollowResponseValidations(rep, senderId, receiverId)
	if err != nil {
		return err
	}

	rep.DeclineFollowRequest(senderId, receiverId)
	return nil
}
