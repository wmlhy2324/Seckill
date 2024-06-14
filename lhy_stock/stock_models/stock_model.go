package stock_models

import "Final_Assessment/common/models"

type StockModel struct {
	models.Model
	ItemId uint  `json:"itemId"` // 商品id
	Stock  int32 `json:"stock"`  // 库存
}
