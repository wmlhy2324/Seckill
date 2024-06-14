package logic

import (
	"Final_Assessment/lhy_cart/cart_models"
	"context"
	"encoding/json"
	"errors"

	"Final_Assessment/lhy_cart/cart_rpc/internal/svc"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCartDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCartDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCartDetailLogic {
	return &GetCartDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 计算购物车内商品的总价和商品数量
func (l *GetCartDetailLogic) GetCartDetail(in *cart_rpc.GetCartDetailRequest) (*cart_rpc.GetCartDetailResponse, error) {
	var item []cart_models.CartItemModel
	var cart cart_models.CartModel
	l.svcCtx.DB.Take(&cart, "user_id = ?", in.UserId)
	l.svcCtx.DB.Where("cart_id = ?", cart.ID).Find(&item)
	if len(item) == 0 {
		err := errors.New("购物车为空")
		return nil, err
	}
	total := 0.0 // 初始化总价为0

	// 遍历购物车中的每一项商品
	for _, cartItem := range item {
		price := cartItem.Price * float64(cartItem.Num) // 计算单项商品总价
		total += price                                  // 累加到总价格上
	}
	data, _ := json.Marshal(item)
	return &cart_rpc.GetCartDetailResponse{
		Price: float32(total),
		Num:   float32(len(item)),
		Data:  data,
	}, nil
}
