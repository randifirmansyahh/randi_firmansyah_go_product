package userRepository

import "randi_firmansyah/app/models/userModel"

type IUserRepository interface {
	FindAll() ([]userModel.User, error)
	FindByID(ID int) (userModel.User, error)
	FindByUsername(username string) (userModel.User, error)
	Create(user userModel.User) (userModel.User, error)
	UpdateV2(user userModel.User) (userModel.User, error)
	Update(id int, user userModel.User) (userModel.User, error)
	Delete(user userModel.User) (userModel.User, error)
}
