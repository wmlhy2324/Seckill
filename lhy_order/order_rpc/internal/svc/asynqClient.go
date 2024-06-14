package svc

import (
	"Final_Assessment/lhy_order/order_rpc/internal/config"
	"github.com/hibiken/asynq"
)

// create asynq client.
func newAsynqClient(c config.Config) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{Addr: c.Redis.Addr, Password: c.Redis.Password})
}
