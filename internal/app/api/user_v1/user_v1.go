package user_v1

import (
	"github.com/iamtonydev/user-service-api/internal/app/service/user"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

type Implementation struct {
	desc.UnimplementedUserV1Server

	userService *user.Service
}

func NewUserV1(userService *user.Service) *Implementation {
	return &Implementation{
		desc.UnimplementedUserV1Server{},

		userService,
	}
}
