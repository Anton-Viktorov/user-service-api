package user_v1

import desc "github.com/iamtonydev/user-service-api/pkg/user_v1"

type Implementation struct {
	desc.UnimplementedUserV1Server
}

func NewUserV1() *Implementation {
	return &Implementation{
		desc.UnimplementedUserV1Server{},
	}
}
