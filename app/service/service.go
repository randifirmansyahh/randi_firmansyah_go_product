package service

import (
	"randi_firmansyah/app/service/cartService"
	"randi_firmansyah/app/service/categoryService"
	"randi_firmansyah/app/service/orderService"
	"randi_firmansyah/app/service/productService"
	"randi_firmansyah/app/service/userService"
)

type Service struct {
	IProductService  productService.IProductService
	IUserService     userService.IUserService
	ICartService     cartService.ICartService
	ICategoryService categoryService.ICategoryService
	IOrderService    orderService.IOrderService
}
