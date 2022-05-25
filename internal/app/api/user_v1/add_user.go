package user_v1

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/convert"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	_ "github.com/jackc/pgx/stdlib"
)

func (i *Implementation) AddUser(ctx context.Context, req *desc.AddUserRequest) (*desc.AddUserResponse, error) {
	id, err := i.userService.AddUser(ctx, convert.ToUserInfo(req))
	if err != nil {
		return nil, err
	}

	return &desc.AddUserResponse{
		Result: &desc.AddUserResponse_Result{
			Id: id,
		},
	}, nil
}
