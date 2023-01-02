package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rinatkh/test_2022/internal/users"
	"github.com/rinatkh/test_2022/internal/users/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/rinatkh/test_2022/pkg/utils"
	"github.com/sirupsen/logrus"
	"strconv"
)

type UserHandler struct {
	userUC users.UseCase
	log    *logrus.Entry
}

func NewUserHandler(userUC users.UseCase, log *logrus.Entry) *UserHandler {
	return &UserHandler{
		userUC: userUC,
		log:    log,
	}
}

func (u UserHandler) DeleteUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.DeleteUserRequest
		params.Id = ctx.Params("user_id")
		data, err := u.userUC.DeleteUser(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u UserHandler) UpdateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.UpdateUserRequest
		params.Id = ctx.Params("user_id")
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.userUC.UpdateUser(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u UserHandler) GetUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetUserRequest
		params.Id = ctx.Params("user_id")

		data, err := u.userUC.GetUser(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u UserHandler) GetUsers() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.GetUsersRequest
		var err error
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
		data, err := u.userUC.GetUsers(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}

func (u UserHandler) CreateUser() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		var params dto.CreateUserRequest
		if err := utils.ReadRequest(ctx, &params); err != nil {
			return constants.InputError
		}

		data, err := u.userUC.CreateUser(&params)
		if err != nil {
			return err
		}
		return ctx.JSON(data)
	}
}
