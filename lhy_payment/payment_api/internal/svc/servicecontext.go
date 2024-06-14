package svc

import (
	"Final_Assessment/lhy_payment/payment_api/internal/config"
	"Final_Assessment/lhy_payment/payment_rpc/payment"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	PaymentRpc payment.Payment
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpc)),
	}
}
