package order_models

import "Final_Assessment/common/models"

type ItemModel struct {
	models.Model
	Name          string  `gorm:"size:32" json:"name"`             // 名称
	Price         float64 `json:"price"`                           // 价格
	Stock         uint    `json:"stock"`                           // 库存
	Description   string  `gorm:"size:128" json:"description"`     // 描述
	IsUseDiscount bool    `json:"isUseDiscount"`                   // 是否使用折扣
	IsDeleted     bool    `gorm:"default:false"  json:"isDeleted"` // 是否删除
}
