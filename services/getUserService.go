package services

import (
	"crudinmemory/customerrs"
	"crudinmemory/repositories"

	"github.com/google/uuid"
)

type GetUserService struct {
	userRepository repositories.UserRepository
}

func NewGetUserService(userRepository repositories.UserRepository) GetUserService {
	return GetUserService{
		userRepository: userRepository,
	}
}

func (c *GetUserService) Execute(id uuid.UUID) (repositories.User, error) {
	user := c.userRepository.FindById(id)

	if (repositories.User{}) == user {
		return repositories.User{}, customerrs.ErrUserNotFoundById
	}

	return user, nil
}
