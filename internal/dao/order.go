package dao

import (
	"trading-system/internal/model"
	"trading-system/pkg/app"
)

func (d *Dao) GetOrderList(page, pageSize int) ([]*model.Order, error) {
	order := model.Order{}
	pageOffset := app.GetPageOffset(page, pageSize)
	return order.List(d.engine, pageOffset, pageSize)
}

func (d *Dao) CreateOrder(owner string, t uint8, quantity, price uint32, price_policy uint8) (model.Order, error) {
	order := model.Order {
		Owner: owner,
		Type: t,
		Quantity: quantity,
		Price: price,
		PricePolicy: price_policy,
	}

	return order.Create(d.engine)
}

func (d *Dao) GetOrder(id uint32) (model.Order, error) {
	order := model.Order {
		ID: id,
	}

	return order.Get(d.engine)
}

func (d *Dao) UpdateOrder(id uint32, owner string, t uint8, quantity, price uint32, price_policy uint8) (model.Order, error) {
	order := model.Order {
		ID: id,
		Owner: owner,
		Type: t,
		Quantity: quantity,
		Price: price,
		PricePolicy: price_policy,
	}

	values := map[string]interface{}{
		"owner":       	owner,
		"type": 		t,
		"quantity":		quantity,
		"price":		price,
		"price_policy":	price_policy,
	}

	return order, order.Update(d.engine, values)
}

func (d *Dao) DeleteOrder(id uint32) error {
	order := model.Order {
		ID: id,
	}

	return order.Delete(d.engine)
}