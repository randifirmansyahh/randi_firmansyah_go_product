package categoryService

import "randi_firmansyah/app/models/categoryModel"

type ICategoryService interface {
	FindAll() ([]categoryModel.Category, error)
	FindByID(id int) (categoryModel.Category, error)
	Create(Category categoryModel.Category) (categoryModel.Category, error)
	Update(id int, Category categoryModel.Category) (categoryModel.Category, error)
	UpdateV2(Category categoryModel.Category) (categoryModel.Category, error)
	Delete(Category categoryModel.Category) (categoryModel.Category, error)
}
