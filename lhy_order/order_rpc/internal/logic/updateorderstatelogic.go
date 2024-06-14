package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_order/order_rpc/internal/svc"
	"Final_Assessment/lhy_order/order_rpc/types/order_rpc"
	"Final_Assessment/mqueue/cmd/job/jobtype"
	"context"
	"encoding/json"
	"errors"
	"github.com/hibiken/asynq"
	Err "github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStateLogic {
	return &UpdateOrderStateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderStateLogic) UpdateOrderState(in *order_rpc.UpdateOrderStateRequest) (*order_rpc.UpdateOrderStateResponse, error) {
	var order *order_models.OrderModel
	l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&order)
	if order.UserId == 0 {
		err := errors.New("orderrpc 订单没有找到")
		return nil, Err.Wrapf(err, "order_id is %v", in.UserId)
	}
	if in.State == order.Status {
		return &order_rpc.UpdateOrderStateResponse{}, nil
	}
	err := l.verifyOrderTradeState(in.State, order.Status)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	order.Status = in.State

	err = l.svcCtx.DB.Save(&order).Error
	if err != nil {
		return nil, errors.New("修改订单状态失败")

	}
	if in.State == order_models.OrderTradeStateWaitPay {
		payload, err := json.Marshal(jobtype.PaySuccessNotifyUserPayload{Order: order})
		if err != nil {
			logx.WithContext(l.ctx).Errorf("notify user fail json.Marshal err:%v", err)
		} else {
			_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.MsgPaySuccessNotifyUser, payload))
			if err != nil {
				logx.WithContext(l.ctx).Errorf("notify user fail asynq.Enqueue err:%v", err)
			}
		}

	}
	return &order_rpc.UpdateOrderStateResponse{
		OrderId:    int32(order.ID),
		State:      order.Status,
		UserId:     int32(order.UserId),
		Price:      float32(order.Amount),
		CreateTime: order.CreatedAt.Unix(),
	}, nil
}

// 验证订单的支付状态
func (l *UpdateOrderStateLogic) verifyOrderTradeState(newTradeState, oldTradeState int64) error {
	if newTradeState == order_models.OrderTradeStateWaitPay {
		return errors.New("订单状态错误,不能修改为待支付")
	}
	if newTradeState == order_models.OrderTradeStateCancel {
		if oldTradeState != order_models.OrderTradeStateWaitPay {
			return errors.New("订单状态错误,只要待支付的订单才能被取消")
		}
	} else if newTradeState == order_models.OrderTradeStateUsed {
		if oldTradeState != order_models.OrderTradeStateWaitPay {
			return errors.New("订单状态错误,只要待支付的订单才能被完成")
		}
	} else if newTradeState == order_models.OrderTradeStateExpire {
		return errors.New("订单已经过期，不能修改状态")
	}
	return nil

}
