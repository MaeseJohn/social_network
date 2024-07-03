package domain

import "time"

type Follow struct {
	FollowerId string
	FollowedId string
	FollowDate string
}

func NewFollow(followerId, followedId string) *Follow {
	date := time.Now().UTC()
	followDate := date.Format(time.DateTime)
	follow := Follow{
		FollowerId: followerId,
		FollowedId: followedId,
		FollowDate: followDate,
	}

	return &follow
}
