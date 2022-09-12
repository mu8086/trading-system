package matching

import (
	"fmt"
	"sort"
	"trading-system/global"
	d "trading-system/internal/dao"
	"trading-system/internal/model"
)

var (
	dao    *d.Dao
	orders []*model.Order
	err    error
)

func initOrders() error {
	dao = d.New(global.DBEngine)
	orders, err = dao.GetOrderList(1, global.AppSetting.DefaultPageSize)
	if err != nil {
		global.Logger.Infof("GetOrderList err: %v", err)
	}

	return err
}

func sortOrders() {
	sort.SliceStable(orders, func(i, j int) bool {
		if orders[i].Type != orders[j].Type {
			return orders[i].Type > orders[j].Type
		}
	
		if orders[i].Type == uint8(model.Buy) {
			if orders[i].PricePolicy == uint8(model.Market) ||
				orders[j].PricePolicy == uint8(model.Market) ||
				orders[i].Price == orders[j].Price {
				return orders[i].ID < orders[j].ID;
			}
			return orders[i].Price < orders[j].Price
		} else if orders[i].Type == uint8(model.Sell) {
			if orders[i].PricePolicy == uint8(model.Market) ||
				orders[j].PricePolicy == uint8(model.Market) ||
				orders[i].Price == orders[j].Price {
				return orders[i].ID > orders[j].ID;
			}
			return orders[i].Price > orders[j].Price
		}
	
		return orders[i].ID < orders[j].ID
	})
}

func printOrders() {
	for k, v := range orders {
		fmt.Printf("[%v]: %v\n", k, *v)
	}
}

func Match() {
	global.Logger.Infof("Match Begin")
	defer func() {
		global.Logger.Infof("Match End")
	}()

	if err := initOrders(); err != nil {
		global.Logger.Fatalf("initOrder error")
		return
	}

	sortOrders()

	printOrders()

	fmt.Printf("workerSize: %v\n", global.MatchSetting.WorkerSize)

}
