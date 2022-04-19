package productService

import (
	"randi_firmansyah/app/models/productModel"
	"randi_firmansyah/app/repository"
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

func (s *service) FindByID(id int) (productModel.Product, error) {
	return s.repository.IProductRepository.FindByID(id)
}

func (s *service) Create(product productModel.Product) (productModel.ProductResponse, error) {
	// convert product to product response
	newProduct, err := s.repository.IProductRepository.Create(product)
	if err != nil {
		return productModel.ProductResponse{}, err
	}

	productResponse := productModel.ProductResponse{
		Id:             newProduct.Id,
		Category_Id:    newProduct.Category_Id,
		Nama:           newProduct.Nama,
		Harga:          newProduct.Harga,
		Qty:            newProduct.Qty,
		Image:          newProduct.Image,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return productResponse, nil
}

func (s *service) Update(id int, product productModel.Product) (productModel.ProductResponse, error) {
	// convert product to product response
	newProduct, err := s.repository.IProductRepository.Update(id, product)
	if err != nil {
		return productModel.ProductResponse{}, err
	}

	productResponse := productModel.ProductResponse{
		Id:             newProduct.Id,
		Category_Id:    newProduct.Category_Id,
		Nama:           newProduct.Nama,
		Harga:          newProduct.Harga,
		Qty:            newProduct.Qty,
		Image:          newProduct.Image,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return productResponse, nil
}

func (s *service) Delete(data productModel.Product) (productModel.ProductResponse, error) {
	newProduct, err := s.repository.IProductRepository.Delete(data)
	if err != nil {
		return productModel.ProductResponse{}, err
	}

	productResponse := productModel.ProductResponse{
		Id:             newProduct.Id,
		Category_Id:    newProduct.Category_Id,
		Nama:           newProduct.Nama,
		Harga:          newProduct.Harga,
		Qty:            newProduct.Qty,
		Image:          newProduct.Image,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return productResponse, nil
}

func (s *service) UpdateV2(product productModel.Product) (productModel.Product, error) {
	return s.repository.IProductRepository.UpdateV2(product)
}
