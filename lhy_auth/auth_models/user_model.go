package auth_models

import "Final_Assessment/common/models"

type UserModel struct {
	models.Model
	Pwd            string `gorm:"size:64" json:"pwd"`
	Nickname       string `gorm:"size:32" json:"nickname"`
	Addr           string `gorm:"size:128" json:"addr"`
	Role           int8   `json:"role"`                          //1为管理员2为普通用户
	OpenID         string `gorm:"size:128" json:"openID"`        //第三方登录凭证
	RegisterSource string `gorm:"size:32" json:"registerSource"` //注册来源
	Account        int32  `json:"account"`                       //账户金额
}
