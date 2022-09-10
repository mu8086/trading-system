package model

import (
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

func (o Order) TableName() string {
	return "trading_order"
}

func (o Order) List(db *gorm.DB, pageOffset, pageSize int) ([]*Order, error) {
	var orders []*Order
	var err error

	if pageOffset >= 0 && pageSize > 0 {
		db = db.Offset(pageOffset).Limit(pageSize)
	}

	if err = db.Where("type != ?", 0).Find(&orders).Error; err != nil {
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