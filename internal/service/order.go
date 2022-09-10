package service

import (
	"trading-system/internal/model"
	"trading-system/pkg/app"
)

type CountOrderRequest struct {
	Owner       string `form:"owner"                    binding:"max=255"`
	Type        uint8  `form:"type"                     binding:"oneof=1 2"`
	Quantity    uint32 `form:"quantity"                 binding:"gt=0"`
	Price       uint32 `form:"price"                    binding:"gt=0"`
	PricePolicy uint8  `form:"price_policy,default=1"   binding:"oneof=1 2"`
}

type OrderListRequest struct {}

type GetOrderRequest struct {
	ID uint32 `form:"id"                                binding:"required,gte=1"`
}

type CreateOrderRequest struct {
	Owner       string `form:"owner"                    binding:"required,min=3,max=255"`
	Type        uint8  `form:"type"                     binding:"required,oneof=1 2"`
	Quantity    uint32 `form:"quantity"                 binding:"required,gt=0"`
	Price       uint32 `form:"price"                    binding:"required,gt=0"`
	PricePolicy uint8  `form:"price_policy,default=1"   binding:"oneof=1 2"`
}

type UpdateOrderRequest struct {
	ID          uint32 `form:"id"                       binding:"required,gte=1"`
	Owner       string `form:"owner"                    binding:"max=255"`
	Type        uint8  `form:"type"                     binding:"oneof=0 1 2"`
	Quantity    uint32 `form:"quantity"                 binding:"gte=0"`
	Price       uint32 `form:"price"                    binding:"gte=0"`
	PricePolicy uint8  `form:"price_policy"             binding:"oneof=0 1 2"`
}

type DeleteOrderRequest struct {
	ID uint32 `form:"id"                                binding:"required,gte=1"`
}

func (svc *Service) GetOrderList(param *OrderListRequest, pager *app.Pager) ([]*model.Order, error) {
	return svc.dao.GetOrderList(pager.Page, pager.PageSize)
}

func (svc *Service) CreateOrder(param *CreateOrderRequest) (model.Order, error) {
	return svc.dao.CreateOrder(param.Owner, param.Type, param.Quantity, param.Price, param.PricePolicy)
}

func (svc *Service) GetOrder(param *GetOrderRequest) (model.Order, error) {
	return svc.dao.GetOrder(param.ID)
}

func (svc *Service) UpdateOrder(param *UpdateOrderRequest) (model.Order, error) {
	return svc.dao.UpdateOrder(param.ID, param.Owner, param.Type, param.Quantity, param.Price, param.PricePolicy)
}

func (svc *Service) DeleteOrder(param *DeleteOrderRequest) error {
	return svc.dao.DeleteOrder(param.ID)
}