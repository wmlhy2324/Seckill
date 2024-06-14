package logic

import (
	"Final_Assessment/lhy_cart/cart_models"
	"Final_Assessment/lhy_cart/cart_rpc/internal/svc"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCartLogic {
	return &CreateCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateCartLogic) CreateCart(in *cart_rpc.CreateCartRequest) (*cart_rpc.CartResponse, error) {
	var cart cart_models.CartModel
	l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&cart)
	if cart.ID == 0 { // 如果记录未找到
		// 创建新的购物车记录
		newCart := cart_models.CartModel{UserId: uint(in.UserId)}
		err := l.svcCtx.DB.Create(&newCart).Error
		if err != nil {
			return nil, err
		}

		return &cart_rpc.CartResponse{
			CartId:  int32(newCart.ID), // 应该返回新创建的购物车的ID
			Message: "购物车创建成功",
		}, nil
	}

	return nil, nil
}
