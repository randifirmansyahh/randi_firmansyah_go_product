package cartService

import (
	"errors"
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

func (s *service) FindByUserID(userId int) ([]cartModel.Cart, error) {
	return s.repository.ICartRepository.FindByUserID(userId)
}

func (s *service) Create(Cart cartModel.Cart) (cartModel.CartResponse, error) {
	findProduct, err := s.repository.IProductRepository.FindByID(Cart.Product_Id)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	Cart.Total = findProduct.Harga * Cart.Qty

	newCart, err := s.repository.ICartRepository.Create(Cart)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	cartResponse := cartModel.CartResponse{
		Id:             newCart.Id,
		Username:       newCart.User.Username,
		Product:        newCart.Product,
		Qty:            newCart.Qty,
		Total:          newCart.Total,
		DateAuditModel: newCart.DateAuditModel,
	}

	return cartResponse, nil
}

func (s *service) Update(id, qty int) error {
	// find cart
	findCart, err := s.repository.ICartRepository.FindByID(id)
	if err != nil {
		return errors.New("cart not found")
	}

	// find product
	findProduct, err := s.repository.IProductRepository.FindByID(findCart.Product_Id)
	if err != nil {
		return errors.New("product not found")
	}

	// check qty
	if qty > findProduct.Qty {
		return errors.New("stok tidak mencukupi")
	}

	// update cart
	findCart.Qty = qty
	findCart.Total = findCart.Product.Harga * findCart.Qty
	_, err = s.repository.ICartRepository.Update(id, findCart)
	if err != nil {
		return errors.New("failed to update cart")
	}

	return nil
}

func (s *service) Delete(data cartModel.Cart) (cartModel.CartResponse, error) {
	newCart, err := s.repository.ICartRepository.Delete(data)
	if err != nil {
		return cartModel.CartResponse{}, err
	}

	cartResponse := cartModel.CartResponse{
		Id:             newCart.Id,
		Username:       newCart.User.Username,
		Product:        newCart.Product,
		Qty:            newCart.Qty,
		Total:          newCart.Total,
		DateAuditModel: newCart.DateAuditModel,
	}
	return cartResponse, nil
}

func (s *service) UpdateV2(Cart cartModel.Cart) (cartModel.Cart, error) {
	return s.repository.ICartRepository.UpdateV2(Cart)
}
