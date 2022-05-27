package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func ToFullUserInfo(req *desc.UpdateUserRequest) *model.UserInfo {
	return &model.UserInfo{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Age:   req.GetAge(),
		Email: req.GetEmail(),
	}
}
