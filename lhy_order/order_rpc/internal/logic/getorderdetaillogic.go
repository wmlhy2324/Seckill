package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_order/order_rpc/internal/svc"
	"Final_Assessment/lhy_order/order_rpc/types/order_rpc"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderDetailLogic {
	return &GetOrderDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderDetailLogic) GetOrderDetail(in *order_rpc.GetOrderDetailRequest) (*order_rpc.GetOrderDetailResponse, error) {
	var order order_models.OrderModel
	err := l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&order).Error
	if err != nil {
		return nil, err
	}

	return &order_rpc.GetOrderDetailResponse{
		OrderId:    int32(order.ID),
		UserId:     in.UserId,
		Price:      float32(order.Amount),
		State:      int64(order.Status),
		CreateTime: order.CreatedAt.Unix(),
	}, nil
}
