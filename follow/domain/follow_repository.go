package domain

type FollowRepository interface {
	SaveFollow(follow *Follow) error
	/*DeleteFollow(follow *Follow) error
	GetFollowers(userId string) error
	GetFollows(userId string) error*/
}
