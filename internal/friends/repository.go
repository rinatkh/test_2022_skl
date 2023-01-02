package friends

type FriendsRepository interface {
	IsFriends(first, second string) (bool, error)
	AddFriend(first, second string) error
	DeleteFriend(first, second string) error
	GetFriends(id string, limit, offset int64) (*[]Friends, int64, error)
}
