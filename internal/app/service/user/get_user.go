package user

import (
	"context"

	"github.com/iamtonydev/user-service-api/internal/app/model"
)

func (s *Service) GetUser(ctx context.Context, userId int64) (*model.UserInfo, error) {
	return s.userRepository.GetUser(ctx, userId)
}
