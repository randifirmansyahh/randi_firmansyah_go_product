package userService

import "randi_firmansyah/app/models/userModel"

type IUserService interface {
	FindAll() ([]userModel.User, error)
}
