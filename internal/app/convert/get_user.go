package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func ToDescGetUserResponse(userInfo *model.UserInfo) *desc.GetUserResponse {
	return &desc.GetUserResponse{
		Result: &desc.GetUserResponse_User{
			Id:    userInfo.Id,
			Name:  userInfo.Name,
			Age:   userInfo.Age,
			Email: userInfo.Email,
		},
	}
}
