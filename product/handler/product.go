package handler

import (
	"context"
	"product/common"
	"product/domain/model"
	"product/domain/service"
	product "product/proto"
)

type Product struct {
	ProductDataService service.IProductDataService
}

// AddProduct 添加商品
func (p *Product) AddProduct(ctx context.Context, req *product.ProductInfo, res *product.ResponseProduct) error {
	productAdd := &model.Product{}
	if err := common.SwapTo(req, productAdd); err != nil {
		return err
	}
	productId, err := p.ProductDataService.AddProduct(productAdd)
	if err != nil {
		return err
	}
	res.ProductId = productId
	return nil
}

// FindProductByID 通过商品id查询商品
func (p *Product) FindProductByID(ctx context.Context, req *product.RequestID, res *product.ProductInfo) error {
	Data, err := p.ProductDataService.FindProductByID(req.Id)
	if err != nil {
		return err
	}
	if err := common.SwapTo(Data, res); err != nil {
		return err
	}
	return err
}

// UpdateProduct 更新商品信息
func (p *Product) UpdateProduct(ctx context.Context, req *product.ProductInfo, res *product.Response) error {
	productInfo := &model.Product{}
	if err := common.SwapTo(req, productInfo); err != nil {
		return err
	}
	if err := p.ProductDataService.UpdateProduct(productInfo); err != nil {
		return err
	}
	res.Msg = "更新成功"
	return nil
}

// DeleteProductByID 根据商品id删除商品
func (p *Product) DeleteProductByID(ctx context.Context, req *product.RequestID, res *product.Response) error {
	err := p.ProductDataService.DeleteProduct(req.Id)
	if err != nil {
		return err
	}
	res.Msg = "删除成功"
	return nil
}
func (p *Product) FindAllProduct(ctx context.Context, req *product.RequestAll, res *product.AllProduct) error {
	productData, err := p.ProductDataService.FindAllProduct()
	if err != nil {
		return err
	}
	if err := common.SwapTo(productData, res.ProductInfo); err != nil {
		return err
	}
	return nil
}
