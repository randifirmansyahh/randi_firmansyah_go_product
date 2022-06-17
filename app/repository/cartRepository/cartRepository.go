package cartRepository

import (
	"randi_firmansyah/app/models/cartModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ICartRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]cartModel.Cart, error) {
	var Carts []cartModel.Cart
	err := r.db.Joins("User").Joins("Product").Find(&Carts).Error
	return Carts, err
}

func (r *repository) FindByID(id int) (cartModel.Cart, error) {
	var Cart cartModel.Cart
	err := r.db.Joins("User").Joins("Product").First(&Cart, id).Error
	return Cart, err
}

func (r *repository) FindByUserID(userId int) ([]cartModel.Cart, error) {
	var Carts []cartModel.Cart
	err := r.db.Joins("User").Joins("Product").Preload("Product.Category").Find(&Carts, "user_id = ?", userId).Error
	return Carts, err
}

func (r *repository) Create(Cart cartModel.Cart) (cartModel.Cart, error) {
	err := r.db.Create(&Cart).Error
	return Cart, err
}

func (r *repository) UpdateV2(Cart cartModel.Cart) (cartModel.Cart, error) {
	err := r.db.Save(&Cart).Error
	return Cart, err
}

func (r *repository) Update(id int, Cart cartModel.Cart) (cartModel.Cart, error) {
	Cart.Id = id
	err := r.db.Model(cartModel.Cart{}).Where("id = ?", id).Updates(Cart).Error
	return Cart, err
}

func (r *repository) Delete(Cart cartModel.Cart) (cartModel.Cart, error) {
	err := r.db.Delete(&Cart).Error
	return Cart, err
}
