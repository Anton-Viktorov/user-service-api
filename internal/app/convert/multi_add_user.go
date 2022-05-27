package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func ToUsersInfo(req *desc.MultiAddUserRequest) []*model.UserInfo {
	users := make([]*model.UserInfo, 0, len(req.GetUsers()))
	for _, user := range req.GetUsers() {
		users = append(users, &model.UserInfo{
			Name:  user.GetName(),
			Age:   user.GetAge(),
			Email: user.GetEmail(),
		})
	}

	return users
}
