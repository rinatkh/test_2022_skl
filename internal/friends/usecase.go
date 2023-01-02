package friends

type UseCase interface {
	AddFriend(params *AddFriendRequest) (*AddFriendResponse, error)
	DeleteFriend(params *DeleteFriendRequest) (*DeleteFriendResponse, error)
	GetFriends(params *GetFriendsRequest) (*GetFriendsResponse, error)
}
