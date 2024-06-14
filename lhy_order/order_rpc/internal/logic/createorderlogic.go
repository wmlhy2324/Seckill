package logic

import (
	"Final_Assessment/lhy_cart/cart_rpc/cart"
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_order/order_rpc/internal/svc"
	"Final_Assessment/lhy_order/order_rpc/types/order_rpc"
	"context"
	"encoding/json"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
	"sync"
)

const CloseOrderTimeMinutes = 1

var writerWg sync.WaitGroup

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrderLogic) CreateOrder(in *order_rpc.CreateOrderRequest) (*order_rpc.CreateOrderResponse, error) {
	var order order_models.OrderModel

	resp, err := l.svcCtx.CartRpc.GetCartDetail(l.ctx, &cart.GetCartDetailRequest{UserId: in.UserId})
	if err != nil {
		return nil, err
	}

	go func() {
		writerWg.Add(1)
		defer writerWg.Done()
		err = l.PushKqOrderSuccess(float64(resp.Price), uint(in.UserId), 0)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("push order fail : %+v", err)

		}
	}()

	writerWg.Wait()
	l.svcCtx.DB.Take(&order, "user_id = ?", in.UserId)
	return &order_rpc.CreateOrderResponse{
		Price:   resp.Price,
		OrderId: int32(order.ID),
	}, nil
}
func (l *CreateOrderLogic) PushKqOrderSuccess(Amount float64, userid uint, status int64) error {
	order := order_models.OrderModel{
		UserId: userid,
		Amount: Amount,
		Status: status,
	}
	body, err := json.Marshal(order)
	if err != nil {
		return errors.New("push kq order fail :json,marshal fail")
	}
	return l.svcCtx.KqueuePutOrderClient.Push(string(body))
}
