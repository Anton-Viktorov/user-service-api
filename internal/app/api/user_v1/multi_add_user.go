package user_v1

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/convert"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) MultiAddUser(ctx context.Context, req *desc.MultiAddUserRequest) (*desc.MultiAddUserResponse, error) {
	count, err := i.userService.MultiAddUser(ctx, convert.ToUsersInfo(req))
	if err != nil {
		return nil, err
	}

	return &desc.MultiAddUserResponse{
		Result: &desc.MultiAddUserResponse_Result{
			Count: count,
		},
	}, nil
}
