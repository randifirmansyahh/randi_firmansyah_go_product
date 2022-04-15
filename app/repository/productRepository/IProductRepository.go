package productRepository

import "randi_firmansyah/app/models/productModel"

type IProductRepository interface {
	FindAll() ([]productModel.Product, error)
	FindByID(ID int) (productModel.Product, error)
	Create(product productModel.Product) (productModel.Product, error)
	UpdateV2(product productModel.Product) (productModel.Product, error)
	Update(id int, product productModel.Product) (productModel.Product, error)
	Delete(product productModel.Product) (productModel.Product, error)
}
