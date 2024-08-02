package domain

import "time"

type FollowRequest struct {
	SenderId    string
	ReceiverId  string
	RequestDate string
	Status      string
}

func NewFollowRequest(senderId, receiverId string) *FollowRequest {
	date := time.Now().UTC()
	followDate := date.Format(time.DateTime)
	followRequest := FollowRequest{
		SenderId:    senderId,
		ReceiverId:  receiverId,
		RequestDate: followDate,
		Status:      "Pending",
	}

	return &followRequest
}
