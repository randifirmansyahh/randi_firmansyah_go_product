package productRepository

import (
	"randi_firmansyah/app/models/productModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IProductRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]productModel.Product, error) {
	var products []productModel.Product
	err := r.db.Joins("Category").Find(&products).Error
	return products, err
}

func (r *repository) FindByID(ID int) (productModel.Product, error) {
	var product productModel.Product
	err := r.db.Joins("Category").First(&product, ID).Error
	return product, err
}

func (r *repository) Create(product productModel.Product) (productModel.Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) UpdateV2(product productModel.Product) (productModel.Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) Update(id int, product productModel.Product) (productModel.Product, error) {
	product.Id = id
	err := r.db.Model(productModel.Product{}).Where("id = ?", id).Updates(product).Error
	return product, err
}

func (r *repository) Delete(product productModel.Product) (productModel.Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
