package cartService

import (
	"randi_firmansyah/app/models/cartModel"
	"randi_firmansyah/app/repository"
)

type service struct {
	repository repository.Repository
}

func NewService(repository repository.Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]cartModel.Cart, error) {
	return s.repository.ICartRepository.FindAll()
}

func (s *service) FindByID(id int) (cartModel.Cart, error) {
	return s.repository.ICartRepository.FindByID(id)
}

func (s *service) Create(Cart cartModel.Cart) (cartModel.CartResponse, error) {
	newCart, err := s.repository.ICartRepository.Create(Cart)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	cartResponse := cartModel.CartResponse{
		Id:             newCart.Id,
		User_Id:        newCart.User_Id,
		Product_Id:     newCart.Product_Id,
		Qty:            newCart.Qty,
		Total:          newCart.Total,
		DateAuditModel: newCart.DateAuditModel,
	}

	return cartResponse, nil
}

func (s *service) Update(id int, Cart cartModel.Cart) (cartModel.CartResponse, error) {
	newCart, err := s.repository.ICartRepository.Update(id, Cart)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	cartResponse := cartModel.CartResponse{
		Id:             newCart.Id,
		User_Id:        newCart.User_Id,
		Product_Id:     newCart.Product_Id,
		Qty:            newCart.Qty,
		Total:          newCart.Total,
		DateAuditModel: newCart.DateAuditModel,
	}

	return cartResponse, nil
}

func (s *service) Delete(data cartModel.Cart) (cartModel.CartResponse, error) {
	newCart, err := s.repository.ICartRepository.Delete(data)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	cartResponse := cartModel.CartResponse{
		Id:             newCart.Id,
		User_Id:        newCart.User_Id,
		Product_Id:     newCart.Product_Id,
		Qty:            newCart.Qty,
		Total:          newCart.Total,
		DateAuditModel: newCart.DateAuditModel,
	}
	return cartResponse, nil
}

func (s *service) UpdateV2(Cart cartModel.Cart) (cartModel.Cart, error) {
	return s.repository.ICartRepository.UpdateV2(Cart)
}
