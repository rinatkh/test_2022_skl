package friends

import "github.com/rinatkh/test_2022/internal/users/models/dto"

type Friends struct {
	FirstUserId  string `json:"first_user" db:"first_user"`
	SecondUserId string `json:"second_user" db:"second_user"`
}

type BasicResponse struct{}

type AddFriendRequest struct {
	Friends
}
type AddFriendResponse struct {
	BasicResponse
}
type DeleteFriendRequest struct {
	Friends
}
type DeleteFriendResponse struct {
	BasicResponse
}
type GetFriendsRequest struct {
	UserId string `query:"user_id"`
	Limit  int64  `query:"limit"`
	Offset int64  `query:"offset"`
}
type GetFriendsResponse struct {
	Friends []dto.User `json:"friends"`
	Length  int64      `json:"length"`
}
