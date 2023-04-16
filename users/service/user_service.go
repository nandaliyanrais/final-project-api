package service

import (
	"mygram-api/models/domain"
	"mygram-api/users/repository"
)

type UserService interface {
	Register(user *domain.User) (err error)
	Login(user *domain.User) (err error)
}

type UserServiceRepository struct {
	UserRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &UserServiceRepository{UserRepository: userRepository}
}

func (userService *UserServiceRepository) Register(user *domain.User) (err error) {
	if err = userService.UserRepository.Register(user); err != nil {
		return err
	}

	return
}

func (userService *UserServiceRepository) Login(user *domain.User) (err error) {
	if err = userService.UserRepository.Login(user); err != nil {
		return err
	}

	return
}