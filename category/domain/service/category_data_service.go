package service

import (
	"category/domain/model"
	"category/domain/repository"
)

type ICategoryDataService interface {
	AddCategory(*model.Category) (int64, error)
	DeleteCategory(int64) error
	UpdateCategory(*model.Category) error
	FindCategoryByID(int64) (*model.Category, error)
	FindAllCategory() ([]model.Category, error)
	FindCategoryByName(string2 string) (*model.Category, error)
	FindCategoryByLevel(uint32) ([]model.Category, error)
	FindCategoryByParent(int64) ([]model.Category, error)
}

func NewCategoryDataService(categoryRepository repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{CategoryRepository: categoryRepository}
}

type CategoryDataService struct {
	CategoryRepository repository.ICategoryRepository
}

// AddCategory 创建分类
func (c *CategoryDataService) AddCategory(category *model.Category) (int64, error) {
	return c.CategoryRepository.CreateCategory(category)
}

// DeleteCategory 删除分类
func (c *CategoryDataService) DeleteCategory(categoryId int64) error {
	return c.CategoryRepository.DeleteCategory(categoryId)
}

// UpdateCategory 更新分类
func (c *CategoryDataService) UpdateCategory(category *model.Category) error {
	return c.CategoryRepository.UpdateCategory(category)
}

// FindCategoryByID 根据ID查找分类
func (c *CategoryDataService) FindCategoryByID(categoryId int64) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByID(categoryId)
}

// FindAllCategory 查找分类
func (c *CategoryDataService) FindAllCategory() ([]model.Category, error) {
	return c.CategoryRepository.FindAllCategory()
}

// FindCategoryByName 根据名字查找分类
func (c *CategoryDataService) FindCategoryByName(name string) (*model.Category, error) {
	return c.CategoryRepository.FindCategoryByName(name)
}

// FindCategoryByLevel 根据等级查找分类
func (c *CategoryDataService) FindCategoryByLevel(level uint32) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByLevel(level)
}

// FindCategoryByParent 根据父类查找分类
func (c *CategoryDataService) FindCategoryByParent(parent int64) ([]model.Category, error) {
	return c.CategoryRepository.FindCategoryByParent(parent)
}
