package user_v1

import (
	"context"
	"fmt"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) RemoveUser(ctx context.Context, req *desc.RemoveUserRequest) (*desc.Empty, error) {
	fmt.Println("=== user deleted ===")
	fmt.Printf("delete user with id: %d\n", req.GetId())

	return &desc.Empty{}, nil
}
