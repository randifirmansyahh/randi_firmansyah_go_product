package categoryRepository

import "randi_firmansyah/app/models/categoryModel"

type ICategoryRepository interface {
	FindAll() ([]categoryModel.Category, error)
	FindByID(id int) (categoryModel.Category, error)
	Create(category categoryModel.Category) (categoryModel.Category, error)
	UpdateV2(category categoryModel.Category) (categoryModel.Category, error)
	Update(id int, Category categoryModel.Category) (categoryModel.Category, error)
	Delete(category categoryModel.Category) (categoryModel.Category, error)
}
