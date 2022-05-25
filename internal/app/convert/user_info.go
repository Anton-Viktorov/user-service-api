package convert

import (
	"github.com/iamtonydev/user-service-api/internal/app/model"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

type UserInfoRequest interface {
	GetName() string
	GetAge() int64
	GetEmail() string
}

type FullUserInfoRequest interface {
	UserInfoRequest
	GetId() int64
}

type UserIdRequest interface {
	GetId() int64
}

func ToUserInfo(req UserInfoRequest) *model.UserInfo {
	return &model.UserInfo{
		Name:  req.GetName(),
		Age:   req.GetAge(),
		Email: req.GetEmail(),
	}
}

func ToFullUserInfo(req FullUserInfoRequest) *model.UserInfo {
	return &model.UserInfo{
		Id:    req.GetId(),
		Name:  req.GetName(),
		Age:   req.GetAge(),
		Email: req.GetEmail(),
	}
}

func ToUserId(req UserIdRequest) int64 {
	return req.GetId()
}

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

func ToUsersInfo(req *desc.MultiAddUserRequest) []*model.UserInfo {
	users := make([]*model.UserInfo, 0)
	for _, user := range req.GetUsers() {
		users = append(users, ToUserInfo(user))
	}

	return users
}

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
