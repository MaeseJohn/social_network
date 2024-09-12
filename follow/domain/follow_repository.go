package domain

type FollowRepository interface {
	RequestFollow(FollowRequest *FollowRequest) error
	SaveFollow(follow *Follow) error
	DeleteFollow(followerid, followedid string) error
	GetUserPrivacity(userid string) (bool, error)
	GetFollowRequests(receiverId string) ([]FollowRequest, error)
	CheckFollowExists(senderId, receiverId string) (bool, error)
	CheckFollowRequestExists(senderId, receiverId string) (bool, error)
	AcceptFollowRequest(follow *Follow, senderId, receiverId string) error
	DeclineFollowRequest(senderId, receiverId string)
	/*GetFollowers(userId string) error
	GetFollows(userId string) error*/
}
