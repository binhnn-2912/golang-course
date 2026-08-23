package service

import "golang-course-api/internal/repository"

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		userRepository: repository.NewUserRepository(),
	}
}

func (us *UserService) GetUsersService() string {
	return us.userRepository.GetUsersRepo()
}
