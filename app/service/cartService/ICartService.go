package cartService

import "randi_firmansyah/app/models/cartModel"

type ICartService interface {
	FindAll() ([]cartModel.Cart, error)
	FindByID(id int) (cartModel.Cart, error)
	FindByUserID(userId int) ([]cartModel.Cart, error)
	Create(cart cartModel.Cart) (cartModel.CartResponse, error)
	Update(id int, qty int) error
	UpdateV2(Cart cartModel.Cart) (cartModel.Cart, error)
	Delete(cart cartModel.Cart) (cartModel.CartResponse, error)
}
