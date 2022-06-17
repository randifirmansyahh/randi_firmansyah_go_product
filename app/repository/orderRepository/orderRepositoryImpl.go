package orderRepository

import (
	"randi_firmansyah/app/models/orderModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IOrderRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]orderModel.Order, error) {
	var orders []orderModel.Order
	err := r.db.Joins("Product").Preload("Product.Category").Find(&orders).Error
	return orders, err
}

func (r *repository) FindByUsername(username string) ([]orderModel.Order, error) {
	var product []orderModel.Order
	err := r.db.Joins("Product").Preload("Product.Category").Find(&product, "username = ?", username).Error
	return product, err
}

func (r *repository) FindByID(ID int) (orderModel.Order, error) {
	var product orderModel.Order
	err := r.db.Joins("Product").Preload("Product.Category").First(&product, ID).Error
	return product, err
}

func (r *repository) Create(product orderModel.Order) (orderModel.Order, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) UpdateV2(product orderModel.Order) (orderModel.Order, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) Update(id int, product orderModel.Order) (orderModel.Order, error) {
	product.Id = id
	err := r.db.Model(orderModel.Order{}).Where("id = ?", id).Updates(product).Error
	return product, err
}

func (r *repository) Delete(product orderModel.Order) (orderModel.Order, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
