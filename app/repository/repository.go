package repository

import (
	"randi_firmansyah/app/repository/productRepository"
	"randi_firmansyah/app/repository/userRepository"
)

type Repository struct {
	IUserRepository    userRepository.IUserRepository
	IProductRepository productRepository.IProductRepository
}
