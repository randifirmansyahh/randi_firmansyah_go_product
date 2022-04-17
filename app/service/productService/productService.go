package productService

import (
	"errors"
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/repository"
	"strconv"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]productModel.Product, error) {
	return s.repository.IProductRepository.FindAll()
}

func (s *service) FindByID(id string) (productModel.Product, error) {
	// check id
	if id == "" {
		return productModel.Product{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return productModel.Product{}, err
	}

	return s.repository.IProductRepository.FindByID(newId)
}

func (s *service) Create(product productModel.Product) (productModel.Product, error) {
	return s.repository.IProductRepository.Create(product)
}

func (s *service) Update(id string, product productModel.Product) (productModel.Product, error) {
	// check id
	if id == "" {
		return productModel.Product{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return productModel.Product{}, err
	}

	return s.repository.IProductRepository.Update(newId, product)
}

func (s *service) Delete(id string) (productModel.Product, error) {
	// check id
	if id == "" {
		return productModel.Product{}, errors.New("id is empty")
	}

	// conv to int
	newId, err := strconv.Atoi(id)
	if err != nil {
		return productModel.Product{}, err
	}

	cari, err := s.repository.IProductRepository.FindByID(newId)
	if err != nil {
		return productModel.Product{}, err
	}

	return s.repository.IProductRepository.Delete(cari)
}

func (s *service) UpdateV2(product productModel.Product) (productModel.Product, error) {
	return s.repository.IProductRepository.UpdateV2(product)
}
