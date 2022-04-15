package service

import (
	"randi_firmansyah/app/service/productService"
	"randi_firmansyah/app/service/userService"
)

type Service struct {
	IProductService productService.IProductService
	IUserService    userService.IUserService
}
