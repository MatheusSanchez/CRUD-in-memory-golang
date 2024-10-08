package services

import "crudinmemory/repositories"

type GetAllUsersService struct {
	userRepository repositories.UserRepository
}

func NewGetAllUsersService(userRepository repositories.UserRepository) GetAllUsersService {
	return GetAllUsersService{
		userRepository: userRepository,
	}
}

func (c *GetAllUsersService) Execute() []repositories.User {

	return c.userRepository.FindAll()
}
