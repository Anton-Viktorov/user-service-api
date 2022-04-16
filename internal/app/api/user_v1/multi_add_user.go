package user_v1

import (
	"context"
	"fmt"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) MultiAddUser(ctx context.Context, req *desc.MultiAddUserRequest) (*desc.MultiAddUserResponse, error) {
	fmt.Println("=== users have been added ===")
	for i, user := range req.GetUsers() {
		fmt.Printf("user %d. name: %s, age: %d, email: %s\n", i, user.GetName(), user.GetAge(), user.GetEmail())
	}

	return &desc.MultiAddUserResponse{
		Result: &desc.MultiAddUserResponse_Result{Id: []int64{1, 2}},
	}, nil
}
