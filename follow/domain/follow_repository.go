package domain

type FollowRepository interface {
	RequestFollow(FollowRequest *FollowRequest) error
	SaveFollow(follow *Follow) error
	DeleteFollow(followerid, followedid string) error
	GetUserPrivacity(userid string) (bool, error)
	GetFollowRequests(receiverId string) ([]FollowRequest, error)
	CheckFriendRequestExists(senderId, receiverId string) (bool, error)
	/*GetFollowers(userId string) error
	GetFollows(userId string) error*/
}
