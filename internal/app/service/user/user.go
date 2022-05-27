package user

import "github.com/iamtonydev/user-service-api/internal/app/repository"

type Service struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}
