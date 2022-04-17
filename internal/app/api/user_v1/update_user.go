package user_v1

import (
	"context"
	"fmt"

	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) UpdateUser(ctx context.Context, req *desc.UpdateUserRequest) (*desc.Empty, error) {
	fmt.Println("=== user updated ===")
	fmt.Printf("id: %d\n", req.GetId())
	fmt.Println("initial data. name: Admin, age: 100500, email: admin@no.no")
	fmt.Printf("update data. name: %s, age: %d, email: %s\n", req.GetName(), req.GetAge(), req.GetEmail())

	return &desc.Empty{}, nil
}
