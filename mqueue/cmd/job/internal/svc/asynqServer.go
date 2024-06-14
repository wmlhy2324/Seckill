package svc

import (
	"Final_Assessment/mqueue/cmd/job/internal/config"
	"fmt"
	"github.com/hibiken/asynq"
)

func newAsynqServer(c config.Config) *asynq.Server {

	return asynq.NewServer(
		asynq.RedisClientOpt{Addr: c.Redis.Addr, Password: c.Redis.Password},
		asynq.Config{
			IsFailure: func(err error) bool {
				fmt.Printf("asynq server exec task IsFailure ======== >>>>>>>>>>>  err : %+v \n", err)
				return true
			},
			Concurrency: 20, //max concurrent process job task num
		},
	)
}
