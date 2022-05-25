package user_v1

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/convert"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
)

func (i *Implementation) GetUser(ctx context.Context, req *desc.GetUserRequest) (*desc.GetUserResponse, error) {
	userInfo, err := i.userService.GetUser(ctx, convert.ToUserId(req))
	if err != nil {
		return nil, err
	}

	return convert.ToDescGetUserResponse(userInfo), nil
}
