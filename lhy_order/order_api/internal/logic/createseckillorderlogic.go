package logic

import (
	"Final_Assessment/lhy_order/order_api/internal/svc"
	"Final_Assessment/lhy_order/order_api/internal/types"
	"Final_Assessment/lhy_order/order_rpc/order"
	"Final_Assessment/lhy_stock/stock_rpc/stock"
	"Final_Assessment/utils/jwts"
	"context"
	"errors"
	"fmt"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

var dtmServer = "etcd://localhost:2379/dtmservice"

type CreateSeckillOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(req *types.CreateSeckillRequest) (resp *types.CreateSeckillOrderResponse, err error) {
	payload, _ := jwts.ParseToken(req.Token, l.svcCtx.Config.Auth.AccessSecret)
	price := l.FindItem(req.ItemId)

	orderRpcServer, err := l.svcCtx.Config.OrderRpc.BuildTarget()
	if err != nil {
		return nil, errors.New("下单异常")
	}
	stockRpcServer, err := l.svcCtx.Config.StockRpc.BuildTarget()
	if err != nil {
		return nil, errors.New("下单异常")
	}
	CreateSeckillReq := &order.CreateSeckillOrderRequest{
		UserId: payload.UserId,
		ItemId: req.ItemId,
	}
	ReduceStockReq := &stock.ReduceInventoryRequest{
		ItemId: req.ItemId,
	}

	gid := dtmgrpc.MustGenGid(dtmServer)

	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRpcServer+"/order_rpc.Order/CreateSeckillOrder", orderRpcServer+"/order_rpc.Order/CreateRollback", CreateSeckillReq).
		Add(stockRpcServer+"/stock_rpc.stock/ReduceInventory", stockRpcServer+"/stock_rpc.stock/ReduceInventoryRollBack", ReduceStockReq)

	err = saga.Submit()

	if err != nil {
		return nil, fmt.Errorf("submit data to  dtm-server err  : %+v \n", err)
	}

	return &types.CreateSeckillOrderResponse{
		Price:  float32(price),
		UserId: int32(payload.UserId),
	}, nil
}

func (l *CreateSeckillOrderLogic) FindItem(ItemId int32) float64 {

	values, err := l.svcCtx.Redis.HGetAll("itemId:" + strconv.Itoa(int(ItemId))).Result()
	if err != nil {
		// 处理错误
		return 0
	}

	priceStr, ok := values["price"]
	if !ok {
		// 价格字段不存在
		return 0
	}

	// 将字符串转换为浮点数
	price, err := strconv.ParseFloat(priceStr, 64)
	if err != nil {
		return 0
	}
	return price
}
