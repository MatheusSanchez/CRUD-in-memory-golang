package services

import (
	"crudinmemory/customerrs"
	"crudinmemory/repositories"

	"github.com/google/uuid"
)

type DeleteUserService struct {
	userRepository repositories.UserRepository
}

func NewDeleteUserService(userRepository repositories.UserRepository) DeleteUserService {
	return DeleteUserService{
		userRepository: userRepository,
	}
}

func (c *DeleteUserService) Execute(id uuid.UUID) (repositories.User, error) {

	deletedUser := c.userRepository.Delete(id)

	if (repositories.User{}) == deletedUser {
		return repositories.User{}, customerrs.ErrUserNotFoundById
	}

	return deletedUser, nil
}
