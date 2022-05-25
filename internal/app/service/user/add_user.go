package user

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/model"
)

func (s *Service) AddUser(ctx context.Context, user *model.UserInfo) (int64, error) {
	return s.userRepository.AddUser(ctx, user)
}
