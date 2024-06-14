package cart_models

import "Final_Assessment/common/models"

type CartItemModel struct {
	models.Model
	Name        string    `gorm:"size:32" json:"name"`         // 名称
	Price       float64   `json:"price"`                       // 价格
	Description string    `gorm:"size:128" json:"description"` // 描述
	CartId      uint      `json:"cartId"`                      // 购物车id
	Num         uint      `json:"num"`                         // 数量
	cart        CartModel `gorm:"foreignKey:CartId"`
}
