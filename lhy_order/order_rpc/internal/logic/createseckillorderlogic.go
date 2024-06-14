package logic

import (
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/mqueue/cmd/job/jobtype"
	"context"
	"encoding/json"
	"errors"
	"github.com/hibiken/asynq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"Final_Assessment/lhy_order/order_rpc/internal/svc"
	"Final_Assessment/lhy_order/order_rpc/types/order_rpc"

	"github.com/zeromicro/go-zero/core/logx"
)

var writer2Wg sync.WaitGroup
var src = rand.NewSource(time.Now().UnixNano())
var rng = rand.New(src)
var minPExpire = 1 * time.Minute
var maxPExpire = 10 * time.Minute
var randomDuration = time.Duration(rng.Int63n(int64(maxPExpire.Minutes()-minPExpire.Minutes()+1))+int64(minPExpire.Minutes())) * time.Minute

type CreateSeckillOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSeckillOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSeckillOrderLogic {
	return &CreateSeckillOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSeckillOrderLogic) CreateSeckillOrder(in *order_rpc.CreateSeckillOrderRequest) (*order_rpc.CreateSeckillOrderResponse, error) {

	var SeckillItem order_models.SeckillItemModel
	cacheKey := "itemId:" + strconv.Itoa(int(in.ItemId))
	itemData, err := l.svcCtx.Redis.HGetAll(cacheKey).Result()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if len(itemData) == 0 {
		l.svcCtx.DB.Take(&SeckillItem, in.ItemId)
		if SeckillItem.ID == 0 {
			return nil, status.Error(codes.Internal, err.Error())
		}
		l.svcCtx.Redis.HMSet(cacheKey, map[string]interface{}{
			"name":        SeckillItem.Name,
			"price":       SeckillItem.Price,
			"stock":       SeckillItem.Stock,
			"description": SeckillItem.Description,
		})
		err = l.svcCtx.Redis.PExpire(cacheKey, randomDuration).Err()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}

	price, err := strconv.ParseFloat(itemData["price"], 64)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
		// 处理转换错误
	}
	go func() {
		writer2Wg.Add(1)
		defer writer2Wg.Done()
		err := l.PushKqOrderSuccess(price, uint(in.UserId), 0)
		if err != nil {
			logx.WithContext(l.ctx).Errorf("push order fail : %+v", err)

		}
	}()
	payload, err := json.Marshal(jobtype.DeferCloseOrderPayload{UserId: int32(in.UserId), ItemId: in.ItemId})
	if err != nil {
		return nil, errors.New("create delay order fail :json,marshal fail")
	} else {
		_, err := l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(jobtype.DeferCloseOrder, payload), asynq.ProcessIn(CloseOrderTimeMinutes*time.Minute))
		if err != nil {
			return nil, errors.New("create delay order fail : insert queue fail")
		}
	}

	if err != nil {
		logx.WithContext(l.ctx).Errorf("ReduceInventory fail : %+v", err)
	}
	writer2Wg.Wait()
	return &order_rpc.CreateSeckillOrderResponse{
		Price: float32(price),
	}, nil
}
func (l *CreateSeckillOrderLogic) PushKqOrderSuccess(Amount float64, userid uint, status int64) error {
	order := order_models.OrderModel{
		UserId: userid,
		Amount: Amount,
		Status: status,
	}
	body, err := json.Marshal(order)
	if err != nil {
		return errors.New("push kq order fail :json,marshal fail")
	}
	return l.svcCtx.KqueuePutOrderClient.Push(string(body))
}
