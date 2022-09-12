package model

import (
	"fmt"

	"gorm.io/gorm"
)

type Order struct {
	ID          uint32 `gorm:"primary_key" json:"id"`
	Owner       string `json:"owner"`
	Type        uint8  `json:"type"` // Buy / Sell
	Quantity    uint32 `json:"quantity"`
	Price       uint32 `json:"price"`
	PricePolicy uint8  `json:"price_policy"` // limit price / market price
}

const (
	Buy		int = iota + 1
	Sell
)

const (
	Limit	int = iota + 1
	Market
)

func (o Order) TableName() string {
	return "trading_order"
}

func (o Order) String() string {
	var t, pricePolicy string

	if o.Type == 1 {
		t = " Buy"
	} else if o.Type == 2 {
		t = "Sell"
	}

	if o.PricePolicy == 1 {
		pricePolicy = "Limit"
	} else if o.PricePolicy == 2 {
		pricePolicy = "Market"
	}

	return fmt.Sprintf("ID: %5v, Owner: %10v, Type: %4v, Quantity: %5v, Price: %5v, PricePolicy: %6v",
		o.ID, o.Owner, t, o.Quantity, o.Price, pricePolicy)
}

func (o Order) List(db *gorm.DB, pageOffset, pageSize int) ([]*Order, error) {
	var orders []*Order
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if err = db.Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (o Order) Create(db *gorm.DB) (Order, error) {
	return o, db.Create(&o).Error
}

func (o Order) Get(db *gorm.DB) (Order, error) {
	var order Order

	db.Where("id = ?", o.ID).Find(&order)
	if order.ID == 0 {
		return order, gorm.ErrRecordNotFound
	}

	if err := db.Error; err != nil {
		return order, err
	}

	return order, nil
}

func (o Order) Update(db *gorm.DB, values interface{}) error {
	return db.Model(&o).Where("id = ?", o.ID).Updates(values).Error
}

func (o Order) Delete(db *gorm.DB) error {
	return db.Where("id = ?", o.ID).Delete(&o).Error
}
