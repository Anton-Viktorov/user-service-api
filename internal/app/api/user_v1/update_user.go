package user_v1

import (
	"context"
	"fmt"

	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.UpdateUserResponse, error) {
	fmt.Println("=== user has been updated ===")
	fmt.Printf("id: %d\n", req.GetId())
	fmt.Println("initial data. name: Admin, age: 100500, email: admin@no.no")
	fmt.Printf("update data. name: %s, age: %d, email: %s\n", req.GetName(), req.GetAge(), req.GetEmail())

	return &desc.UpdateUserResponse{
		Result: &desc.UpdateUserResponse_User{
			Name:  req.GetName(),
			Age:   req.GetAge(),
			Email: req.GetEmail(),
		},
	}, nil
}
