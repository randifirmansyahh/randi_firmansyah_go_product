package categoryService

import (
	"randi_firmansyah/app/models/categoryModel"
	"randi_firmansyah/app/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]categoryModel.Category, error) {
	return s.repository.ICategoryRepository.FindAll()
}

func (s *service) FindByID(id int) (categoryModel.Category, error) {
	return s.repository.ICategoryRepository.FindByID(id)
}

func (s *service) Create(Category categoryModel.Category) (categoryModel.Category, error) {
	return s.repository.ICategoryRepository.Create(Category)
}

func (s *service) Update(id int, Category categoryModel.Category) (categoryModel.Category, error) {
	return s.repository.ICategoryRepository.Update(id, Category)
}

func (s *service) Delete(data categoryModel.Category) (categoryModel.Category, error) {
	return s.repository.ICategoryRepository.Delete(data)
}

func (s *service) UpdateV2(Category categoryModel.Category) (categoryModel.Category, error) {
	return s.repository.ICategoryRepository.UpdateV2(Category)
}
