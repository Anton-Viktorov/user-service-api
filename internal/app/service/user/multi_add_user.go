package user

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/model"
)

func (s *Service) MultiAddUser(ctx context.Context, users []*model.UserInfo) (int64, error) {
	return s.userRepository.MultiAddUser(ctx, users)
}
