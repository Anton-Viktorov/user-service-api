package user

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/model"
)

func (s *Service) UpdateUser(ctx context.Context, user *model.UserInfo) error {
	return s.userRepository.UpdateUser(ctx, user)
}
