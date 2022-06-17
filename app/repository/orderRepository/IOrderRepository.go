package orderRepository

import "randi_firmansyah/app/models/orderModel"

type IOrderRepository interface {
	FindAll() ([]orderModel.Order, error)
	FindByUsername(username string) ([]orderModel.Order, error)
	FindByID(ID int) (orderModel.Order, error)
	Create(product orderModel.Order) (orderModel.Order, error)
	UpdateV2(product orderModel.Order) (orderModel.Order, error)
	Update(id int, product orderModel.Order) (orderModel.Order, error)
	Delete(product orderModel.Order) (orderModel.Order, error)
}
