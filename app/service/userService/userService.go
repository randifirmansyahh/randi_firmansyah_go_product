package userService

import (
	"errors"
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
	data, err := s.repository.IUserRepository.FindByUsername(username)
	if err != nil {
		return userModel.User{}, errors.New("user tidak ditemukan")
	}

	return data, nil
}

func (s *service) Create(user userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Create(user)
}

func (s *service) Update(id int, User userModel.User) (userModel.User, error) {
	data, err := s.repository.IUserRepository.Update(id, User)
	if err != nil {
		return userModel.User{}, errors.New("data user tidak berhasil diubah")
	}
	return data, nil
}

func (s *service) UpdateV2(user userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.UpdateV2(user)
}

func (s *service) Delete(data userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.Delete(data)
}
