package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func ToUserInfo(req *desc.AddUserRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:  req.GetName(),
		Age:   req.GetAge(),
		Email: req.GetEmail(),
	}
}
