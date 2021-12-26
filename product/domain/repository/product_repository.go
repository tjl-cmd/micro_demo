package repository

import (
	"github.com/jinzhu/gorm"
	"product/domain/model"
)

type IProductRepository interface {
	InitTable() error
	DeleteProductByID(int64) error
	CreateProduct(*model.Product) (int64, error)
	FindProductByID(int64) (*model.Product, error)
	UpdateProduct(*model.Product) error
	FindAll() ([]model.Product, error)
}

type ProductRepository struct {
	mysqlDb *gorm.DB
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{mysqlDb: db}
}
func (p *ProductRepository) InitTable() error {
	return p.mysqlDb.CreateTable(&model.Product{}, &model.ProductImage{}, &model.ProductSize{}, &model.ProductSeo{}).Error
}
func (p *ProductRepository) FindProductByID(productID int64) (product *model.Product, err error) {
	product = &model.Product{}
	return product, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").First(product, productID).Error
}
func (p *ProductRepository) CreateProduct(product *model.Product) (int64, error) {
	return product.ID, p.mysqlDb.Create(product).Error
}
func (p *ProductRepository) DeleteProductByID(ProductID int64) error {
	tx := p.mysqlDb.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if tx.Error != nil {
		return tx.Error
	}
	// 删除
	if err := tx.Unscoped().Where("id = ?", ProductID).Delete(&model.Product{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("images_product_id  = ?", ProductID).Delete(&model.ProductImage{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("size_product_id = ?", ProductID).Delete(&model.ProductSize{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Unscoped().Where("seo_product_id = ?", ProductID).Delete(&model.ProductSeo{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error

}
func (p *ProductRepository) UpdateProduct(product *model.Product) error {
	return p.mysqlDb.Model(product).Update(product).Error
}
func (p *ProductRepository) FindAll() (productAll []model.Product, err error) {
	return productAll, p.mysqlDb.Preload("ProductImage").Preload("ProductSize").Preload("ProductSeo").Find(&productAll).Error
}
