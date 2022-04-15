package userService

import (
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]userModel.User, error) {
	return s.repository.IUserRepository.FindAll()
}
