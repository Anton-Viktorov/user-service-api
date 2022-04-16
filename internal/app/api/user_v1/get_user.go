package user_v1

import (
	"context"
	"fmt"

	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	fmt.Println("=== user has been requested ===")
	fmt.Printf("user id: %d\n", req.GetId())

	return &desc.GetUserResponse{
		Result: &desc.GetUserResponse_User{
			Name:  "Admin",
			Age:   100500,
			Email: "admin@no.no",
		},
	}, nil
}
