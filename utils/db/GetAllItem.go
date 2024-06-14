package db

import (
	"Final_Assessment/core"
	"Final_Assessment/utils/GetGormFields"
	"Final_Assessment/utils/db/config"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type Item struct {
	ID    uint    `gorm:"column:id;primaryKey"`
	Name  string  `column:"name"`  // 名称
	Price float64 `column:"price"` // 价格
	Stock uint    `column:"stock"` // 库存
}

var (
	_all_item_field = GetGormFields.GetGormFields(Item{})
)
var configFile = flag.String("ff", "etc/cartitem.yaml", "the config file")

func GetAllItem(ch chan<- Item) {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	MysqlBD := core.InitGorm(c.Mysql.Datasource)
	const PAGE_SIZE = 10
	var maxid uint

	maxid = 0
	for {
		var Items []Item
		err := MysqlBD.Select(_all_item_field).Where("id>?", maxid).Limit(PAGE_SIZE).Find(&Items).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				logx.Info("read table %s failed: %s", Item{}.TableName(), err)
			}
			break
		}
		if len(Items) == 0 {
			break
		}
		for _, Item := range Items {
			if Item.ID > maxid {
				maxid = Item.ID
			}
			ch <- Item
		}
	}
	close(ch) //close后就不允许再往channel里添加元素了
}
func (Item) TableName() string {
	return "seckill_item_models"
}
