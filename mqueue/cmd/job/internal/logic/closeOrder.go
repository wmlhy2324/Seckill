package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_order/order_rpc/order"
	"Final_Assessment/mqueue/cmd/job/internal/svc"
	"Final_Assessment/mqueue/cmd/job/jobtype"
	"context"
	"encoding/json"
	"errors"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

//var ErrCloseOrderFal = xerr.NewErrMsg("close order fail")

// CloseHomestayOrderHandler close no pay homestayOrder
type CloseOrderHandler struct {
	svcCtx *svc.ServiceContext
}

func NewCloseOrderHandler(svcCtx *svc.ServiceContext) *CloseOrderHandler {
	return &CloseOrderHandler{
		svcCtx: svcCtx,
	}
}
func (l *CloseOrderHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var p jobtype.DeferCloseOrderPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return errors.New("close order fail json.Unmarshal err")
	}
	resp, err := l.svcCtx.OrderRpc.GetOrderDetail(ctx, &order.GetOrderDetailRequest{
		UserId: p.UserId,
	})
	if resp.UserId == 0 || err != nil {
		return errors.New("close order fail GetOrderDetail err")
	}
	if resp.State == order_models.OrderTradeStateWaitPay {
		_, err := l.svcCtx.OrderRpc.UpdateOrderState(ctx, &order.UpdateOrderStateRequest{
			UserId: resp.UserId,
			State:  order_models.OrderTradeStateCancel,
		})
		if err != nil {
			return errors.New("close order fail UpdateOrderState err")
		}
	}
	//l.svcCtx.CartRpc.IncreaseInventory(ctx,&cart.IncreaseInventoryRequest{ItemId: p.ItemId})
	logx.Info("close order success")
	return nil
}
