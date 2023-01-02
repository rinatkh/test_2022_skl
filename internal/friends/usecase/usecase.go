package usecase

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/friends"
	"github.com/rinatkh/test_2022/internal/users"
	"github.com/rinatkh/test_2022/internal/users/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
)

type FriendsUseCase struct {
	cfg         *config.Config
	log         *logrus.Entry
	repoFriends friends.FriendsRepository
	userUC      users.UseCase
}

func NewFriendsUC(cfg *config.Config, log *logrus.Entry, repoFriends friends.FriendsRepository, userUC users.UseCase) friends.UseCase {
	return &FriendsUseCase{
		cfg:         cfg,
		log:         log,
		repoFriends: repoFriends,
		userUC:      userUC,
	}
}

func (u FriendsUseCase) AddFriend(params *friends.AddFriendRequest) (*friends.AddFriendResponse, error) {
	isFriends, err := u.repoFriends.IsFriends(params.FirstUserId, params.SecondUserId)
	if err != nil {
		return nil, err
	}
	if isFriends {
		return nil, constants.NewCodedError("You are already friends", fiber.StatusConflict)
	}
	err = u.repoFriends.AddFriend(params.FirstUserId, params.SecondUserId)
	if err != nil {
		return nil, err
	}
	return &friends.AddFriendResponse{}, nil
}

func (u FriendsUseCase) DeleteFriend(params *friends.DeleteFriendRequest) (*friends.DeleteFriendResponse, error) {
	isFriends, err := u.repoFriends.IsFriends(params.FirstUserId, params.SecondUserId)
	if err != nil {
		return nil, err
	}
	if !isFriends {
		return nil, constants.NewCodedError("You are not already friends", fiber.StatusConflict)
	}
	err = u.repoFriends.DeleteFriend(params.FirstUserId, params.SecondUserId)
	if err != nil {
		return nil, err
	}
	return &friends.DeleteFriendResponse{}, nil
}

func (u FriendsUseCase) GetFriends(params *friends.GetFriendsRequest) (*friends.GetFriendsResponse, error) {
	list, length, err := u.repoFriends.GetFriends(params.UserId, params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	var result []dto.User
	var friendID string
	for _, i := range *list {
		if i.FirstUserId == params.UserId {
			friendID = i.FirstUserId
		} else {
			friendID = i.SecondUserId
		}
		user, err := u.userUC.GetUser(&dto.GetUserRequest{Id: friendID})
		if err != nil {
			return nil, constants.NewCodedError(fmt.Sprintf("User '%s' not exist", friendID), fiber.StatusConflict)
		}
		result = append(result, user.User)
	}
	return &friends.GetFriendsResponse{
		Friends: result,
		Length:  length,
	}, nil
}
