package repository

import (
	"randi_firmansyah/app/repository/cartRepository"
	"randi_firmansyah/app/repository/categoryRepository"
	"randi_firmansyah/app/repository/orderRepository"
	"randi_firmansyah/app/repository/productRepository"
	"randi_firmansyah/app/repository/userRepository"
)

type Repository struct {
	IUserRepository     userRepository.IUserRepository
	IProductRepository  productRepository.IProductRepository
	ICartRepository     cartRepository.ICartRepository
	ICategoryRepository categoryRepository.ICategoryRepository
	IOrderRepository    orderRepository.IOrderRepository
}
