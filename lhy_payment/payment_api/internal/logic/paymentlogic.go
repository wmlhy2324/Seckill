package logic

import (
	"Final_Assessment/lhy_payment/payment_rpc/types/payment_rpc"
	"context"

	"Final_Assessment/lhy_payment/payment_api/internal/svc"
	"Final_Assessment/lhy_payment/payment_api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaymentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaymentLogic {
	return &PaymentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PaymentLogic) Payment(req *types.PayRequest) (resp *types.PayResponse, err error) {
	res, err := l.svcCtx.PaymentRpc.DeductAccount(l.ctx, &payment_rpc.DeductAccountRequest{
		UserId:  req.UserId,
		OrderId: req.OrderId,
	})
	if err != nil {
		return nil, err

	}

	return &types.PayResponse{
		Account: float64(res.Account),
		Msg:     "支付成功",
	}, nil
}
