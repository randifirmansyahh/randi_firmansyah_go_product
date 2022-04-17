package userService

import (
	"errors"
	"randi_firmansyah/app/helper/helper"
	"randi_firmansyah/app/models/userModel"
	"randi_firmansyah/app/repository"
	"strconv"
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

func (s *service) FindByID(id string) (userModel.User, error) {
	// check id
	if id == "" {
		return userModel.User{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return userModel.User{}, err
	}
	return s.repository.IUserRepository.FindByID(newId)
}

func (s *service) FindByUsername(username string) (userModel.User, error) {
	// check id
	if username == "" {
		return userModel.User{}, errors.New("id is empty")
	}

	return s.repository.IUserRepository.FindByUsername(username)
}

func (s *service) Create(user userModel.User) (userModel.User, error) {
	newPassword := helper.Encode([]byte(user.Password))
	user.Password = string(newPassword)
	return s.repository.IUserRepository.Create(user)
}

func (s *service) Update(id string, User userModel.User) (userModel.User, error) {
	// check id
	if id == "" {
		return userModel.User{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return userModel.User{}, err
	}

	return s.repository.IUserRepository.Update(newId, User)
}

func (s *service) UpdateV2(user userModel.User) (userModel.User, error) {
	return s.repository.IUserRepository.UpdateV2(user)
}

func (s *service) Delete(id string) (userModel.User, error) {
	// check id
	if id == "" {
		return userModel.User{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return userModel.User{}, err
	}

	cari, err := s.repository.IUserRepository.FindByID(newId)
	if err != nil {
		return userModel.User{}, err
	}

	return s.repository.IUserRepository.Delete(cari)
}
