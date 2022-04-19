package userService

import "randi_firmansyah/app/models/userModel"

type IUserService interface {
	FindAll() ([]userModel.User, error)
	FindByID(id int) (userModel.User, error)
	FindByUsername(username string) (userModel.User, error)
	Create(user userModel.User) (userModel.User, error)
	Update(id int, User userModel.User) (userModel.User, error)
	UpdateV2(user userModel.User) (userModel.User, error)
	Delete(userModel.User) (userModel.User, error)
}
