package productService

import "randi_firmansyah/app/models/productModel"

type IProductService interface {
	FindAll() ([]productModel.Product, error)
	FindByID(id string) (productModel.Product, error)
	Create(product productModel.Product) (productModel.Product, error)
	Update(id string, product productModel.Product) (productModel.Product, error)
	// UpdateV2(id, string, product productModel.Product) (productModel.Product, error)
	Delete(id string) (productModel.Product, error)
}
