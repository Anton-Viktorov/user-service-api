package user_v1

import (
	"context"
	"fmt"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) ListUser(ctx context.Context, req *desc.Empty) (*desc.ListUserResponse, error) {
	fmt.Println("=== users have been requested ===")

	return &desc.ListUserResponse{
		Result: &desc.ListUserResponse_Result{
			User: []*desc.ListUserResponse_Result_User{
				{
					Name:  "Admin",
					Age:   100500,
					Email: "admin@no.no",
				},
				{
					Name:  "Admin2",
					Age:   100501,
					Email: "admin2@no.no",
				},
			},
		},
	}, nil
}
