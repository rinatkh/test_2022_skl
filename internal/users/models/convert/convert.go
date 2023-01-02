package convert

import (
	"github.com/rinatkh/test_2022/internal/users/models/core"
	"github.com/rinatkh/test_2022/internal/users/models/dto"
	"github.com/rinatkh/test_2022/pkg/utils"
	"time"
)

func User2DTO(user *core.User) dto.User {
	result := dto.User{
		Id:        user.Id,
		Firstname: user.Firstname,
		Surname:   user.Surname,
		Age:       utils.RoundTime(time.Now().Sub(user.BirthDate).Seconds() / 31207680),
		Sex:       user.Sex,
	}
	if user.Middlename != nil {
		result.Middlename = *user.Middlename
	}
	return result
}
