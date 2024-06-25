package domain

import "time"

type Follow struct {
	FollowerId string
	FollowedId string
	FollowDate string
}

func NewFollow(followerId, followedId string) *Follow {
	time := time.Now().UTC()
	followDate := time.Format("2024-06-25 14:30:00")
	follow := Follow{
		FollowerId: followerId,
		FollowedId: followedId,
		FollowDate: followDate,
	}

	return &follow
}
