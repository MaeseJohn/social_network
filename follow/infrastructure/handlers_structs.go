package infrastructure

type followedid struct {
	FollowedId string `validate:"required,uuid"`
}
