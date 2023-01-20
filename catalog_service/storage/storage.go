package storage

import e "github.com/uacademy/e_commerce/catalog_service/proto-gen/ecommerce"

type StorageI interface {
	CreateProduct(id string, entity *e.CreateProductRequest) error
	GetProductList(offset, limit int, search string) (resp *e.GetProductListResponse, err error)
	UpdateProduct(entity *e.UpdateProductRequest) error
	GetProductById(id string) (resp *e.GetProductByIdResponse, err error)
	DeleteProduct(id string) error

	CreateCategory(id string, entity *e.CreateCategoryRequest) error
	GetCategoryList(offset, limit int, search string) (resp *e.GetCategoryListResponse, err error)
	UpdateCategory(entity *e.UpdateCategoryRequest) error
	GetCategoryById(id string) (resp *e.GetCategoryByIdResponse, err error)
	DeleteCategory(id string) error
}
