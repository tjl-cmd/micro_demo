package handler

import (
	"cart/domain/model"
	"cart/domain/service"
	cart "cart/proto"
	"context"
	"micro_demo/common"
)

type Cart struct{
	CateDatService service.ICartDataService
}


func (c *Cart)AddCart(ctx context.Context, req *cart.CartInfo, res *cart.ResponseAdd) error{
	cart :=model.Cart{}
	common.SwapTo()
}
func (c *Cart)CleanCart(ctx context.Context, req * cart.Clean, res *cart.Response) error{

}
func (c *Cart)Incr(ctx context.Context, req *cart.Item, res *cart.Response) error{

}
func (c *Cart)Decr(ctx context.Context, req *cart.Item, res *cart.Response) error{

}
func (c *Cart)DeleteItemByID(ctx context.Context, req *cart.CartID, res *cart.Response) error{

}
func (c *Cart)GetAll(ctx context.Context, req *cart.CartFindAll, res *cart.CartAll) error{

}