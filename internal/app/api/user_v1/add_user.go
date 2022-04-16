package user_v1

import (
	"context"
	"fmt"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	fmt.Println("=== user has been added ===")
	fmt.Printf("name: %s, age: %d, email: %s\n", req.GetName(), req.GetAge(), req.GetEmail())

	return &desc.AddUserResponse{
		Result: &desc.AddUserResponse_Result{
			Id: 1,
		},
	}, nil
}
