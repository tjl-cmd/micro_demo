package repository

import (
	"category/domain/model"
	"github.com/jinzhu/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	CreateCategory(*model.Category) (int64, error)
	UpdateCategory(*model.Category) error
	DeleteCategory(int64) error
	FindCategoryByName(string) (*model.Category, error)
	FindCategoryByID(int64) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
	FindAllCategory() ([]model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{mysqlDb: db}
}

type CategoryRepository struct {
	mysqlDb *gorm.DB
}

func (c *CategoryRepository) InitTable() error {
	return c.mysqlDb.CreateTable(&model.Category{}).Error
}
func (c *CategoryRepository) FindCategoryByID(categoryId int64) (category *model.Category, err error) {
	return category, c.mysqlDb.Where("id = ?", categoryId).Find(category).Error
}
func (c *CategoryRepository) CreateCategory(category *model.Category) (int64, error) {
	return category.ID, c.mysqlDb.CreateTable(category).Error
}
func (c *CategoryRepository) DeleteCategory(categoryId int64) error {
	return c.mysqlDb.Where("id = ?", categoryId).Delete(&model.Category{}).Error
}
func (c *CategoryRepository) UpdateCategory(category *model.Category) error {
	return c.mysqlDb.Model(category).Update(category).Error
}
func (c *CategoryRepository) FindAllCategory() (category []model.Category, err error) {
	return category, c.mysqlDb.Find(&category).Error
}
func (c *CategoryRepository) FindCategoryByName(categoryName string) (category *model.Category, err error) {
	category = &model.Category{}
	return category, c.mysqlDb.Where("category_name = ?", categoryName).Find(category).Error
}
func (c *CategoryRepository) FindCategoryByLevel(level uint32) (category []model.Category, err error) {
	return category, c.mysqlDb.Where("category_level = ?", level).Find(&category).Error
}
func (c *CategoryRepository) FindCategoryByParent(parent int64) (category []model.Category, err error) {
	return category, c.mysqlDb.Where("category_parent = ?", parent).Find(category).Error
}
