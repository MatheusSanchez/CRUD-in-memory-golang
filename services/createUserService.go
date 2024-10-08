package services

import "crudinmemory/repositories"

type CreateUserService struct {
	userRepository repositories.UserRepository
}

func NewCreateUserService(userRepository repositories.UserRepository) CreateUserService {
	return CreateUserService{
		userRepository: userRepository,
	}
}

func (c *CreateUserService) Execute(firstName string, lastName string, bio string) repositories.User {

	return c.userRepository.Insert(firstName, lastName, bio)
}
