package user_v1

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/convert"
	desc "github.com/iamtonydev/user-service-api/pkg/user_v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) ListUser(ctx context.Context, req *emptypb.Empty) (*desc.ListUserResponse, error) {
	usersInfo, err := i.userService.ListUser(ctx)
	if err != nil {
		return nil, err
	}

	return convert.ToDescListUserResponse(usersInfo), nil
}
