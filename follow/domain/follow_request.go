package domain

import (
	"time"

	"github.com/google/uuid"
)

type FollowRequest struct {
	FollowId    string
	SenderId    string
	ReceiverId  string
	Status      string
	RequestDate string
}

func NewFollowRequest(senderId, receiverId string) *FollowRequest {
	date := time.Now().UTC()
	followDate := date.Format(time.DateTime)
	followRequest := FollowRequest{
		FollowId:    uuid.NewString(),
		SenderId:    senderId,
		ReceiverId:  receiverId,
		Status:      "Pending",
		RequestDate: followDate,
	}

	return &followRequest
}
