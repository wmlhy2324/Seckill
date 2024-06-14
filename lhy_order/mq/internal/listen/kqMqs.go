package listen

import (
	"Final_Assessment/lhy_order/mq/internal/config"
	kqMq "Final_Assessment/lhy_order/mq/internal/mqs/kq"
	"Final_Assessment/lhy_order/mq/internal/svc"
	"context"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

func KqMqs(c config.Config, ctx context.Context, svcContext *svc.ServiceContext) []service.Service {

	return []service.Service{
		//Listening for changes in consumption flow status
		kq.MustNewQueue(c.PutOrderConf, kqMq.NewPutOrder(ctx, svcContext)),
		//.....
	}

}
