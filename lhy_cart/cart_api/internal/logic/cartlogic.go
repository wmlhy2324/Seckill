package logic

import (
	"Final_Assessment/lhy_cart/cart_api/internal/svc"
	"Final_Assessment/lhy_cart/cart_api/internal/types"
	"Final_Assessment/lhy_cart/cart_models"
	"Final_Assessment/lhy_cart/cart_rpc/types/cart_rpc"
	"Final_Assessment/utils/jwts"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type CartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CartLogic {
	return &CartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CartLogic) Cart(req *types.CreateCartRequest) (resp *types.CreateCartResponse, err error) {
	var cart cart_models.CartModel
	var item cart_models.ItemModel
	payload, err := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	if err != nil {
		return nil, err

	}
	_, err = l.svcCtx.CartRpc.CreateCart(context.Background(), &cart_rpc.CreateCartRequest{
		UserId: int32(payload.UserId),
	})

	l.svcCtx.DB.Where("name = ?", req.ItemName).First(&item)
	l.svcCtx.DB.Where("user_id = ?", payload.UserId).First(&cart)
	_, err = l.svcCtx.CartRpc.AddItem(context.Background(), &cart_rpc.AddItemRequest{
		CartId:      int32(cart.ID),
		ItemName:    req.ItemName,
		Num:         int32(req.Num),
		Price:       int32(item.Price),
		Description: item.Description,
	})
	if err != nil {
		err = errors.New("加入购物车失败,商品不存在或库存不足")
		return nil, err
	}

	return &types.CreateCartResponse{
		Msg: "加入购物车成功",
	}, nil
}
