package orderService

import (
	"randi_firmansyah/app/models/orderModel"
)

type IOrderService interface {
	FindAll() ([]orderModel.Order, error)
	FindByUsername(username string) ([]orderModel.Order, error)
	FindByID(id int) (orderModel.Order, error)
	Create(product orderModel.Order) (orderModel.OrderResponse, error)
	Update(id int, product orderModel.Order) (orderModel.OrderResponse, error)
	UpdateV2(product orderModel.Order) (orderModel.Order, error)
	Delete(data orderModel.Order) (orderModel.OrderResponse, error)
}
