package logic

import (
	"Final_Assessment/lhy_order/order_rpc/order"
	"Final_Assessment/utils/jwts"
	"context"
	"errors"

	"Final_Assessment/lhy_order/order_api/internal/svc"
	"Final_Assessment/lhy_order/order_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateOrderLogic) CreateOrder(req *types.CreateOrderRequest) (resp *types.CreateOrderResponse, err error) {
	payload, _ := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	res, err := l.svcCtx.OrderRpc.CreateOrder(l.ctx, &order.CreateOrderRequest{
		UserId: int32(payload.UserId),
	})
	if err != nil {
		return nil, errors.New("创建订单失败")
	}

	return &types.CreateOrderResponse{
		Price:   res.Price,
		OrderId: res.OrderId,
		UserId:  int32(payload.UserId),
	}, nil
}
