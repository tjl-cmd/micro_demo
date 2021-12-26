package service

import (
	"product/domain/model"
	"product/domain/repository"
)

type IProductDataService interface {
	AddProduct(*model.Product) (int64, error)
	DeleteProduct(int64) error
	UpdateProduct(*model.Product) error
	FindProductByID(int64) (*model.Product, error)
	FindAllProduct() ([]model.Product, error)
}

type ProductDataService struct {
	ProductDataService repository.IProductRepository
}

func NewProductDataService(productRepository repository.IProductRepository) IProductDataService {
	return &ProductDataService{productRepository}
}
func (p *ProductDataService) AddProduct(product *model.Product) (int64, error) {
	return p.ProductDataService.CreateProduct(product)

}
func (p *ProductDataService) DeleteProduct(productId int64) error {
	return p.ProductDataService.DeleteProductByID(productId)
}
func (p *ProductDataService) UpdateProduct(product *model.Product) error {
	return p.ProductDataService.UpdateProduct(product)
}
func (p *ProductDataService) FindProductByID(productId int64) (*model.Product, error) {
	return p.ProductDataService.FindProductByID(productId)

}
func (p *ProductDataService) FindAllProduct() ([]model.Product, error) {
	return p.ProductDataService.FindAll()
}
