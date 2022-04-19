package productService

import "randi_firmansyah/app/models/productModel"

type IProductService interface {
	FindAll() ([]productModel.Product, error)
	FindByID(id int) (productModel.Product, error)
	Create(product productModel.Product) (productModel.ProductResponse, error)
	Update(id int, product productModel.Product) (productModel.ProductResponse, error)
	UpdateV2(product productModel.Product) (productModel.Product, error)
	Delete(data productModel.Product) (productModel.ProductResponse, error)
}
