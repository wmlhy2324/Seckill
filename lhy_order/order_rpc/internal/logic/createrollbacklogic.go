package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_order/order_rpc/internal/svc"
	"Final_Assessment/lhy_order/order_rpc/types/order_rpc"
	"context"
	"database/sql"
	"fmt"

	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRollbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRollbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRollbackLogic {
	return &CreateRollbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRollbackLogic) CreateRollback(in *order_rpc.CreateSeckillOrderRequest) (*order_rpc.CreateSeckillOrderResponse, error) {
	var order order_models.OrderModel

	err := l.svcCtx.DB.Where("user_id = ?", in.UserId).Find(&order).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if order.ID != 0 {
		barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
		if err != nil {

			return nil, status.Error(codes.Internal, err.Error())
		}
		err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
			query := "DELETE FROM order_models WHERE `user_id` = ?"
			_, err := tx.Exec(query, in.UserId)
			if err != nil {
				return fmt.Errorf("回滚订单失败  err : %v , userId:%d", err, in.UserId)
			}
			return nil
		})
		if err != nil {
			logx.Info(err)
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &order_rpc.CreateSeckillOrderResponse{}, nil
}
