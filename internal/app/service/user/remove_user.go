package user

import (
	"context"
)

func (s *Service) RemoveUser(ctx context.Context, userId int64) error {
	return s.userRepository.RemoveUser(ctx, userId)
}
