package usecase

import (
	"github.com/golang/mock/gomock"
	"github.com/rinatkh/test_2022/internal/test"
	"github.com/rinatkh/test_2022/internal/users"
	mock_users "github.com/rinatkh/test_2022/internal/users/mocks"
	"github.com/rinatkh/test_2022/internal/users/models/core"
	"github.com/rinatkh/test_2022/internal/users/models/dto"
	"github.com/rinatkh/test_2022/pkg/constants"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	testRepo := mock_users.NewMockUserRepository(ctrl)
	dbUserImpl := NewUserUC(test.TestConfig(t), test.TestLogger(t), testRepo)

	type Input struct {
		Limit  int64
		Offset int64
	}
	type InputGetUsers struct {
		Limit  int64
		Offset int64
	}
	type OutputGetUsers struct {
		Users  *[]core.User
		Length int64
		Err    error
	}
	type Output struct {
		Res *dto.GetUsersResponse
		Err error
	}

	str := "1999-11-18"
	date, _ := time.Parse("2006-01-02", str)

	tests := []struct {
		name           string
		input          Input
		inputGetUsers  InputGetUsers
		outputGetUsers OutputGetUsers
		output         Output
	}{
		{
			name: "Don't found in BD",
			input: Input{
				Limit:  0,
				Offset: 0,
			},
			inputGetUsers: InputGetUsers{
				Limit:  0,
				Offset: 0,
			},
			outputGetUsers: OutputGetUsers{
				Users:  nil,
				Length: 0,
				Err:    constants.ErrDB,
			},
			output: Output{
				Res: nil,
				Err: constants.ErrDB,
			},
		},
		{
			name: "Success",
			input: Input{
				Limit:  5,
				Offset: 0,
			},
			inputGetUsers: InputGetUsers{
				Limit:  5,
				Offset: 0,
			},
			outputGetUsers: OutputGetUsers{
				Users: &[]core.User{{
					Id:         "1234567890",
					Firstname:  "Ivan",
					Surname:    "Ivanov",
					Middlename: nil,
					Fio:        "Ivan Ivanov",
					Sex:        "m",
					BirthDate:  date,
				}},
				Length: 1,
				Err:    nil,
			},
			output: Output{
				Res: &dto.GetUsersResponse{
					Users: []dto.User{{
						Id:         "1234567890",
						Firstname:  "Ivan",
						Surname:    "Ivanov",
						Middlename: "",
						Age:        23,
						Sex:        "m",
					}},
					Length: 1,
				},

				Err: nil,
			},
		},
	}
	gomock.InOrder(
		testRepo.EXPECT().GetUsers(tests[0].inputGetUsers.Limit, tests[0].inputGetUsers.Offset).Return(tests[0].outputGetUsers.Users, tests[0].outputGetUsers.Length, tests[0].outputGetUsers.Err),
		testRepo.EXPECT().GetUsers(tests[1].inputGetUsers.Limit, tests[1].inputGetUsers.Offset).Return(tests[1].outputGetUsers.Users, tests[1].outputGetUsers.Length, tests[1].outputGetUsers.Err),
	)
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			res, errRes := users.UseCase.GetUsers(dbUserImpl, &dto.GetUsersRequest{Limit: test.input.Limit, Offset: test.input.Offset})
			if !assert.Equal(t, test.output.Res, res) {
				t.Error("got : ", res, " expected :", test.output.Res)
			}
			if !assert.Equal(t, test.output.Err, errRes) {
				t.Error("got : ", errRes, " expected :", test.output.Err)
			}
		})
	}
}
