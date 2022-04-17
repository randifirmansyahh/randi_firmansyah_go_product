package userService

import "randi_firmansyah/app/models/userModel"

type IUserService interface {
	FindAll() ([]userModel.User, error)
	FindByID(id string) (userModel.User, error)
	Create(user userModel.User) (userModel.User, error)
	Update(id string, User userModel.User) (userModel.User, error)
	UpdateV2(user userModel.User) (userModel.User, error)
	Delete(id string) (userModel.User, error)
}
