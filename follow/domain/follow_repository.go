package domain

type FollowRepository interface {
	SaveFollow(follow *Follow) error
	DeleteFollow(followerid, followedid string) error
	/*GetFollowers(userId string) error
	GetFollows(userId string) error*/
}
