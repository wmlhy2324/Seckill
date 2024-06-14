package cart_models

import "Final_Assessment/common/models"

type CartModel struct {
	models.Model
	UserId uint      `json:"userId"` // 用户id
	User   UserModel `gorm:"foreignKey:UserId;references:ID"`
}
