package userService

import (
	"randi_firmansyah/app/helper/helper"
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

func (s *service) FindByID(id int) (userModel.User, error) {
	return s.repository.IUserRepository.FindByID(id)
}

func (s *service) FindByUsername(username string) (userModel.User, error) {
	return s.repository.IUserRepository.FindByUsername(username)
}

func (s *service) Create(user userModel.User) (userModel.User, error) {
	newPassword := helper.Encode([]byte(user.Password))
	user.Password = string(newPassword)
	return s.repository.IUserRepository.Create(user)
}

func (s *service) Update(id int, User userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Update(id, User)
}

func (s *service) UpdateV2(user userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.UpdateV2(user)
}

func (s *service) Delete(data userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Delete(data)
}
