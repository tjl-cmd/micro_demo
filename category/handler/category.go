package handler

import (
	"category/common"
	"category/domain/model"
	"category/domain/service"
	pb "category/proto"
	"context"
	"encoding/json"
)

type Category struct {
	CategoryDataService service.ICategoryDataService
}

// CreateCategory 创建分类
func (c *Category) CreateCategory(ctx context.Context, in *pb.CategoryRequest, response *pb.CreateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(in, category)
	if err != nil {
		return err
	}
	categoryId, err := c.CategoryDataService.AddCategory(category)
	if err != nil {
		return err
	}
	response.CategoryId = categoryId
	return nil
}

// UpdateCategory 更新分类
func (c *Category) UpdateCategory(ctx context.Context, in *pb.CategoryRequest, response *pb.UpdateCategoryResponse) error {
	category := &model.Category{}
	err := common.SwapTo(in, category)
	if err != nil {
		return err
	}
	err = c.CategoryDataService.UpdateCategory(category)
	if err != nil {
		return err
	}
	response.Message = "分类更新成功"
	return nil
}

// DeleteCategory 删除分类
func (c *Category) DeleteCategory(ctx context.Context, in *pb.DeleteCategoryRequest, response *pb.DeleteCategoryResponse) error {
	err := c.CategoryDataService.DeleteCategory(in.CategoryId)
	if err != nil {
		return err
	}
	response.Message = "分类删除成功"
	return nil
}

// FindCategoryByName 根据名字查询分类
func (c *Category) FindCategoryByName(ctx context.Context, in *pb.FindByNameRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByName(in.CategoryName)
	if err != nil {
		return err
	}
	response = CategoryForResponse(category)

	return nil
}

// FindCategoryByID 根据id查询分类信息
func (c *Category) FindCategoryByID(ctx context.Context, in *pb.FindByIdRequest, response *pb.CategoryResponse) error {
	category, err := c.CategoryDataService.FindCategoryByID(in.CategoryId)
	if err != nil {
		return err
	}
	response = CategoryForResponse(category)
	return nil
}

// FindCategoryByLevel 根据level查询分类
func (c *Category) FindCategoryByLevel(ctx context.Context, in *pb.FindByLevelRequest, response *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindCategoryByLevel(in.Level)
	if err != nil {
		return err
	}
	response, err = CategoryForResponses(category)

	if err != nil {
		return err
	}
	return nil
}

// FindCategoryByParent 根据父类查询分类
func (c *Category) FindCategoryByParent(ctx context.Context, in *pb.FindByParentRequest, response *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindCategoryByParent(in.ParentId)
	if err != nil {
		return err
	}
	response, err = CategoryForResponses(category)
	if err != nil {
		return err
	}
	return nil
}

// FindAllCategory 查询所有分类
func (c *Category) FindAllCategory(ctx context.Context, in *pb.FindAllRequest, response *pb.FindAllResponse) error {
	category, err := c.CategoryDataService.FindAllCategory()
	if err != nil {
		return err
	}
	response, err = CategoryForResponses(category)
	if err != nil {
		return err
	}
	return nil
}
func CategoryForResponse(category *model.Category) (response *pb.CategoryResponse) {
	response.CategoryName = category.CategoryName
	response.CategoryLevel = category.CategoryLevel
	response.CategoryDescription = category.CategoryDescription
	response.CategoryParent = category.CategoryParent
	response.Id = category.ID
	return
}
func CategoryForResponses(category []model.Category) (response *pb.FindAllResponse, err error) {
	categoryByte, err := json.Marshal(category)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(categoryByte, &response.Category)
	if err != nil {
		return nil, err
	}
	return response, nil
}
