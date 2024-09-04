package infrastructure

type followedid struct {
	ReceiverId string `validate:"required,uuid"`
}
