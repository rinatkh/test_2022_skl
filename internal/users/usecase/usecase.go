package usecase

import (
	"github.com/rinatkh/test_2022/config"
	"github.com/rinatkh/test_2022/internal/users"
	"github.com/rinatkh/test_2022/internal/users/models/convert"
	"github.com/rinatkh/test_2022/internal/users/models/core"
	"github.com/rinatkh/test_2022/internal/users/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/sirupsen/logrus"
	"time"
)

type UserUseCase struct {
	cfg      *config.Config
	log      *logrus.Entry
	repoUser users.UserRepository
}

func NewUserUC(cfg *config.Config, log *logrus.Entry, repoUser users.UserRepository) users.UseCase {
	return &UserUseCase{
		cfg:      cfg,
		log:      log,
		repoUser: repoUser,
	}
}

func (u UserUseCase) CreateUser(params *dto.CreateUserRequest) (*dto.CreateUserResponse, error) {
	date, err := time.Parse("2006-01-02", params.BirthDate)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	user := core.User{
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: &params.Middlename,
		Sex:        params.Sex,
		BirthDate:  date,
	}
	result, err := u.repoUser.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dto.CreateUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) UpdateUser(params *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error) {
	check, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrUserDBNotFound
	}
	date, err := time.Parse("2006-01-02", params.BirthDate)
	if err != nil {
		return nil, constants.ErrConvertData
	}
	user := core.User{
		Id:         params.Id,
		Firstname:  params.Firstname,
		Surname:    params.Surname,
		Middlename: &params.Middlename,
		Sex:        params.Sex,
		BirthDate:  date,
	}
	result, err := u.repoUser.UpdateUser(&user)
	if err != nil {
		return nil, err
	}

	return &dto.UpdateUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) DeleteUser(params *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error) {
	check, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if check == nil {
		return nil, constants.ErrUserDBNotFound
	}
	err = u.repoUser.DeleteUser(params.Id)
	if err != nil {
		return nil, err
	}
	return &dto.DeleteUserResponse{}, nil
}

func (u UserUseCase) GetUser(params *dto.GetUserRequest) (*dto.GetUserResponse, error) {
	result, err := u.repoUser.GetUserById(params.Id)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, constants.ErrUserDBNotFound
	}
	return &dto.GetUserResponse{User: convert.User2DTO(result)}, nil
}

func (u UserUseCase) GetUsers(params *dto.GetUsersRequest) (*dto.GetUsersResponse, error) {
	list, length, err := u.repoUser.GetUsers(params.Limit, params.Offset)
	if err != nil {
		return nil, err
	}
	if list == nil {
		return &dto.GetUsersResponse{}, nil
	}
	var result []dto.User
	for _, i := range *list {
		result = append(result, convert.User2DTO(&i))
	}
	return &dto.GetUsersResponse{Users: result, Length: length}, nil
}
