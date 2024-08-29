package domain

import "time"

type FollowRequest struct {
	SenderId    string
	ReceiverId  string
	Status      string
	RequestDate string
}

func NewFollowRequest(senderId, receiverId string) *FollowRequest {
	date := time.Now().UTC()
	followDate := date.Format(time.DateTime)
	followRequest := FollowRequest{
		SenderId:    senderId,
		ReceiverId:  receiverId,
		Status:      "Pending",
		RequestDate: followDate,
	}

	return &followRequest
}
