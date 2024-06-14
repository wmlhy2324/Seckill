package logic

import (
	"Final_Assessment/lhy_cart/cart_models"
	"context"
	"errors"

	"Final_Assessment/lhy_cart/cart_rpc/internal/svc"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *cart_rpc.AddItemRequest) (*cart_rpc.AddItemResponse, error) {
	var cartitem cart_models.CartItemModel
	var item cart_models.ItemModel
	l.svcCtx.DB.Where("name = ?", in.ItemName).Find(&item)
	if item.ID == 0 {
		err := errors.New("商品不存在")
		return nil, err
	}
	if uint(in.Num) > item.Stock {
		err := errors.New("库存不足")
		return nil, err
	}
	l.svcCtx.DB.Where("cart_id = ? and name = ?", in.CartId, in.ItemName).FirstOrCreate(&cartitem)
	cartitem.Name = in.ItemName
	cartitem.CartId = uint(in.CartId)
	cartitem.Num = uint(in.Num)
	cartitem.Price = float64(in.Price)
	cartitem.Description = in.Description
	err := l.svcCtx.DB.Save(&cartitem).Error
	if err != nil {
		return nil, err

	}
	return &cart_rpc.AddItemResponse{
		CartId:  in.CartId,
		Message: "加入购物车成功",
	}, nil
}
