package userRepository

import (
	"randi_firmansyah/app/models/userModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IUserRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]userModel.User, error) {
	var users []userModel.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *repository) FindByID(id int) (userModel.User, error) {
	var user userModel.User
	err := r.db.First(&user, id).Error
	return user, err
}

func (r *repository) FindByUsername(username string) (userModel.User, error) {
	var user userModel.User
	err := r.db.First(&user, "username = ?", username).Error
	return user, err
}

func (r *repository) Create(user userModel.User) (userModel.User, error) {
	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) UpdateV2(user userModel.User) (userModel.User, error) {
	err := r.db.Save(&user).Error
	return user, err
}

func (r *repository) Update(id int, user userModel.User) (userModel.User, error) {
	err := r.db.Model(userModel.User{}).Where("id = ?", id).Updates(user).Error
	return user, err
}

func (r *repository) Delete(user userModel.User) (userModel.User, error) {
	err := r.db.Delete(&user).Error
	return user, err
}
