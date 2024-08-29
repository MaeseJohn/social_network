package infrastructure

import (
	"fmt"
	"social_media/db"
	"social_media/follow/domain"
)

type PostgresRepository struct {
}

func NewPostgresRepository() *PostgresRepository {
	return &PostgresRepository{}
}

func (*PostgresRepository) SaveFollow(follow *domain.Follow) error {
	_, err := db.DataBase().
		NamedExec("INSERT INTO follows (follower_id, followed_id, follow_date) VALUES (:followerid, :followedid, :followdate)", follow)

	if err != nil {
		fmt.Println(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (*PostgresRepository) DeleteFollow(followerid, followedid string) error {

	_, err := db.DataBase().
		Exec("DELETE FROM follows WHERE follower_id=$1 AND followed_id=$2", followerid, followedid)

	if err != nil {
		fmt.Println(err)
		return domain.ErrInternalServerError
	}

	return nil
}

func (*PostgresRepository) GetUserPrivacity(userid string) (bool, error) {
	var private []bool
	err := db.DataBase().
		Select(&private, "SELECT private FROM users WHERE user_id=$1", userid)
	if err != nil {
		return false, domain.ErrNotFound
	}

	return private[0], nil
}

//////////////////
//Follow request//
//////////////////

func (*PostgresRepository) RequestFollow(followRequest *domain.FollowRequest) error {
	_, err := db.DataBase().
		NamedExec("INSERT INTO follow_requests (sender_id, receiver_id, request_date, status) VALUES (:senderid, :receiverid, :requestdate, :status)", followRequest)

	if err != nil {
		return domain.ErrInternalServerError
	}

	return nil
}

func (*PostgresRepository) CheckFriendRequestExists(senderId, receiverId string) (bool, error) {
	var exists bool
	err := db.DataBase().Get(&exists, "SELECT EXISTS(SELECT 1 FROM follow_requests WHERE sender_id = $1 AND receiver_id = $2)", senderId, receiverId)
	if err != nil {
		return false, domain.ErrInternalServerError
	}

	return exists, nil
}

type DBFollowRequest struct {
	SenderId    string `db:"sender_id"`
	RequestDate string `db:"request_date"`
}

func (followRequest *DBFollowRequest) DBFollowRequestToFollowRequest() domain.FollowRequest {
	return domain.FollowRequest{
		SenderId:    followRequest.SenderId,
		ReceiverId:  "",
		Status:      "Pending",
		RequestDate: followRequest.RequestDate,
	}
}

func (*PostgresRepository) GetFollowRequests(receiverid string) ([]domain.FollowRequest, error) {
	var DBFollowRequests []DBFollowRequest
	err := db.DataBase().
		Select(&DBFollowRequests, "SELECT sender_id, request_date FROM follow_requests WHERE receiver_id=$1 AND status=$2", receiverid, "Pending")
	if err != nil {
		return nil, domain.ErrNotFound
	}
	var followRequests []domain.FollowRequest
	for _, request := range DBFollowRequests {
		followRequests = append(followRequests, request.DBFollowRequestToFollowRequest())
	}

	return followRequests, nil
}