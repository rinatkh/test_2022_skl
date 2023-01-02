package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/friends"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type FriendHandler struct {
	friendUC friends.UseCase
	log      *logrus.Entry
}

func NewFriendHandler(friendUC friends.UseCase, log *logrus.Entry) *FriendHandler {
	return &FriendHandler{
		friendUC: friendUC,
		log:      log,
	}
}

func (u FriendHandler) DeleteFriend() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params friends.DeleteFriendRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.friendUC.DeleteFriend(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u FriendHandler) AddFriend() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params friends.AddFriendRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.friendUC.AddFriend(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u FriendHandler) GetFriends() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params friends.GetFriendsRequest
		var err error
		params.UserId = ctx.Query("user_id")
		params.Limit, err = strconv.ParseInt(ctx.Query("limit", "20"),
			10, 64)
		if err != nil {
			return err
		}
		params.Offset, err = strconv.ParseInt(ctx.Query("offset", "0"),
			10, 64)
		if err != nil {
			return err
		}

		data, err := u.friendUC.GetFriends(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
