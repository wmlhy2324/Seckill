package main

import (
	"Final_Assessment/common/barrier"
	"Final_Assessment/core"
	"Final_Assessment/lhy_cart/cart_models"
	"Final_Assessment/lhy_discount/discount_models"
	"Final_Assessment/lhy_order/order_models"
	"Final_Assessment/lhy_stock/stock_models"
	"Final_Assessment/lhy_user/user_models"
	"flag"
	"fmt"
)

type Options struct {
	DB bool
}

func main() {
	var opt Options
	flag.BoolVar(&opt.DB, "db", false, "db")
	flag.Parse()
	if opt.DB {
		db := core.InitGorm("root:112304@tcp(127.0.0.1:3306)/seckill_db?charset=utf8mb4&parseTime=True&loc=Local")
		err := db.AutoMigrate(&user_models.UserModel{},

			&order_models.OrderModel{},
			&discount_models.DiscountModel{},
			&cart_models.ItemModel{},
			&cart_models.CartModel{},
			&cart_models.CartItemModel{},
			&order_models.SeckillItemModel{},
			&barrier.BarrierModel{},
			&stock_models.StockModel{},
		)
		if err != nil {
			fmt.Println("表结构创建失败", err)
			return
		}
		fmt.Println("表结构创建成功")
	}
	fmt.Println("初始化成功")
}
