package cartRepository

import "randi_firmansyah/app/models/cartModel"

type ICartRepository interface {
	FindAll() ([]cartModel.Cart, error)
	FindByID(id int) (cartModel.Cart, error)
	Create(cart cartModel.Cart) (cartModel.Cart, error)
	UpdateV2(cart cartModel.Cart) (cartModel.Cart, error)
	Update(id int, Cart cartModel.Cart) (cartModel.Cart, error)
	Delete(cart cartModel.Cart) (cartModel.Cart, error)
}
