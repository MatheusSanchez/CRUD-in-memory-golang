package services

import (
	"crudinmemory/repositories"

	"github.com/google/uuid"
)

type EditUserService struct {
	userRepository repositories.UserRepository
}

func NewEditUserService(userRepository repositories.UserRepository) EditUserService {
	return EditUserService{
		userRepository: userRepository,
	}
}

func (c *EditUserService) Execute(id uuid.UUID, firstName string, lastName string, bio string) repositories.User {

	return c.userRepository.Edit(id, firstName, lastName, bio)
}
