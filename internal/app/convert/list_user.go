package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func ToDescListUserResponse(usersInfo []*model.UserInfo) *desc.ListUserResponse {
	users := make([]*desc.ListUserResponse_Result_User, 0, len(usersInfo))

	for _, userInfo := range usersInfo {
		users = append(
			users,
			&desc.ListUserResponse_Result_User{
				Id:    userInfo.Id,
				Name:  userInfo.Name,
				Age:   userInfo.Age,
				Email: userInfo.Email,
			},
		)
	}

	return &desc.ListUserResponse{
		Result: &desc.ListUserResponse_Result{
			Users: users,
		},
	}
}
