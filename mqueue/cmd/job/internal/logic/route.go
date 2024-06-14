package logic

import (
	"Final_Assessment/mqueue/cmd/job/internal/svc"
	"Final_Assessment/mqueue/cmd/job/jobtype"
	"context"
	"github.com/hibiken/asynq"
)

type CronJob struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCronJob(ctx context.Context, svcCtx *svc.ServiceContext) *CronJob {
	return &CronJob{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// register job
func (l *CronJob) Register() *asynq.ServeMux {

	mux := asynq.NewServeMux()

	//defer job
	mux.Handle(jobtype.DeferCloseOrder, NewCloseOrderHandler(l.svcCtx))

	return mux
}
