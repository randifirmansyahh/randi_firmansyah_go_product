package orderService

import (
	"errors"
	"randi_firmansyah/app/models/orderModel"
	"randi_firmansyah/app/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]orderModel.Order, error) {
	return s.repository.IOrderRepository.FindAll()
}

func (s *service) FindByUsername(username string) ([]orderModel.Order, error) {
	// find user
	if _, err := s.repository.IUserRepository.FindByUsername(username); err != nil {
		return nil, errors.New("user tidak ditemukan")
	}

	// find order
	data, err := s.repository.IOrderRepository.FindByUsername(username)
	if err != nil {
		return nil, errors.New("user belum memiliki order")
	}

	return data, nil
}

func (s *service) FindByID(id int) (orderModel.Order, error) {
	data, err := s.repository.IOrderRepository.FindByID(id)
	if err != nil {
		return orderModel.Order{}, errors.New("order tidak ditemukan")
	}
	return data, nil
}

func (s *service) Create(product orderModel.Order) (orderModel.OrderResponse, error) {
	// convert order to order response
	newProduct, err := s.repository.IOrderRepository.Create(product)
	if err != nil {
		return orderModel.OrderResponse{}, err
	}

	orderResponse := orderModel.OrderResponse{
		Id:             newProduct.Id,
		Username:       newProduct.Username,
		Product:        newProduct.Product,
		Qty:            newProduct.Qty,
		Total:          newProduct.Total,
		OrderStatus:    newProduct.OrderStatus,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return orderResponse, nil
}

func (s *service) Update(id int, product orderModel.Order) (orderModel.OrderResponse, error) {
	// convert order to order response
	newProduct, err := s.repository.IOrderRepository.Update(id, product)
	if err != nil {
		return orderModel.OrderResponse{}, err
	}

	orderResponse := orderModel.OrderResponse{
		Id:             newProduct.Id,
		Username:       newProduct.Username,
		Product:        newProduct.Product,
		Qty:            newProduct.Qty,
		Total:          newProduct.Total,
		OrderStatus:    newProduct.OrderStatus,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return orderResponse, nil
}

func (s *service) Delete(data orderModel.Order) (orderModel.OrderResponse, error) {
	newProduct, err := s.repository.IOrderRepository.Delete(data)
	if err != nil {
		return orderModel.OrderResponse{}, err
	}

	orderResponse := orderModel.OrderResponse{
		Id:             newProduct.Id,
		Username:       newProduct.Username,
		Product:        newProduct.Product,
		Qty:            newProduct.Qty,
		Total:          newProduct.Total,
		OrderStatus:    newProduct.OrderStatus,
		DateAuditModel: newProduct.DateAuditModel,
	}

	return orderResponse, nil
}

func (s *service) UpdateV2(product orderModel.Order) (orderModel.Order, error) {
	return s.repository.IOrderRepository.UpdateV2(product)
}
