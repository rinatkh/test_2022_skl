package users

import "github.com/rinatkh/test_2022/internal/users/models/dto"

type UseCase interface {
	CreateUser(params *dto.CreateUserRequest) (*dto.CreateUserResponse, error)
	UpdateUser(params *dto.UpdateUserRequest) (*dto.UpdateUserResponse, error)
	DeleteUser(params *dto.DeleteUserRequest) (*dto.DeleteUserResponse, error)
	GetUser(params *dto.GetUserRequest) (*dto.GetUserResponse, error)
	GetUsers(params *dto.GetUsersRequest) (*dto.GetUsersResponse, error)
}
