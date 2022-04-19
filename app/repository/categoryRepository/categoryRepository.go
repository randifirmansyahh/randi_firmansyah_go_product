package categoryRepository

import (
	"randi_firmansyah/app/models/categoryModel"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) ICategoryRepository {
	return &repository{db: db}
}

func (r *repository) FindAll() ([]categoryModel.Category, error) {
	var Category []categoryModel.Category
	err := r.db.Find(&Category).Error
	return Category, err
}

func (r *repository) FindByID(id int) (categoryModel.Category, error) {
	var Category categoryModel.Category
	err := r.db.First(&Category, id).Error
	return Category, err
}

func (r *repository) Create(Category categoryModel.Category) (categoryModel.Category, error) {
	err := r.db.Create(&Category).Error
	return Category, err
}

func (r *repository) UpdateV2(Category categoryModel.Category) (categoryModel.Category, error) {
	err := r.db.Save(&Category).Error
	return Category, err
}

func (r *repository) Update(id int, Category categoryModel.Category) (categoryModel.Category, error) {
	Category.Id = id
	err := r.db.Model(categoryModel.Category{}).Where("id = ?", id).Updates(Category).Error
	return Category, err
}

func (r *repository) Delete(Category categoryModel.Category) (categoryModel.Category, error) {
	err := r.db.Delete(&Category).Error
	return Category, err
}
