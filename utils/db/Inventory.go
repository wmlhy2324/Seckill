package db

import (
	"Final_Assessment/core"
	"Final_Assessment/utils/db/config"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

const (
	prefix = "item_count_" //所有key设置统一的前缀，方便后续按前缀遍历key
)

func InitItemInventory() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	itemCh := make(chan Item, 100)
	go GetAllItem(itemCh)

	client := core.InitRedis(c.Redis.Addr, c.Redis.Password, c.Redis.DB)

	for {
		item, ok := <-itemCh
		if !ok { //channel已经消费完了
			break
		}
		if item.Stock <= 0 {

			continue //没有库存的商品不参与
		}
		err := client.Set(prefix+strconv.Itoa(int(item.ID)), item.Stock, 0).Err() //0表示不设过期时间
		if err != nil {
			logx.Info("set item %d:%s count to %d failed: %s", item.ID, item.Name, item.Price, err)
		}
	}

}
