package order_models

// -1 取消订单 0 待支付 1 已支付 2 已过期
var OrderTradeStateCancel int64 = -1
var OrderTradeStateWaitPay int64 = 0
var OrderTradeStateUsed int64 = 1
var OrderTradeStateExpire int64 = 2
