package user

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/model"
)

func (s *Service) ListUser(ctx context.Context) ([]*model.UserInfo, error) {
	return s.userRepository.ListUser(ctx)
}
