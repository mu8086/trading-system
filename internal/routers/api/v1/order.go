package v1

import (
	"net/http"
	"trading-system/global"
	"trading-system/internal/service"
	"trading-system/pkg/app"
	"trading-system/pkg/convert"
	"trading-system/pkg/errcode"

	"github.com/gin-gonic/gin"
)

type Order struct{}

func NewOrder() Order {
	return Order{}
}

// @Summary 取得多個訂單
// @Produce json
// @Success 200 {object} model.Order 	"成功"
// @Failure 400 {object} errcode.Error 	"請求錯誤"
// @Failure 500 {object} errcode.Error 	"內部錯誤"
// @Router /api/v1/orders [get]
func (o Order) List(c *gin.Context) {
	//app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	//c.JSON(200, "List")
	param := service.OrderListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	orders, err := svc.GetOrderList(&param, &pager)
	if err != nil {
		global.Logger.Infof("svc.GetOrderList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderListFail)
		return
	}

	response.ToResponseList(orders, len(orders))
	//response.ToResponse(gin.H{})
}

// @Summary 新增訂單
// @Produce json
// @Param owner 		query 	string 	true	"訂單擁有者"	minlength(3) 	maxlength(255)
// @Param type 			query 	int 	true	"訂單種類, 1為買, 2為賣" 		Enums(1, 2)
// @Param quantity 		query 	int 	true	"交易數量"
// @Param price 		query 	int 	true	"交易價格"
// @Param price_policy 	query 	int 	false	"價格原則, 1為限價, 2為市價" 		Enums(1, 2) 	default(1)
// @Success 201 {object} model.Order 	"成功"
// @Failure 400 {object} errcode.Error 	"請求錯誤"
// @Failure 500 {object} errcode.Error 	"內部錯誤"
// @Router /api/v1/orders [post]
func (o Order) Create(c *gin.Context) {
	param := service.CreateOrderRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	order, err := svc.CreateOrder(&param)
	if err != nil {
		global.Logger.Infof("svc.CreateOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateOrderFail)
		return
	}

	response.ToResponseAndStatus(order, http.StatusCreated)
}

// @Summary 取得單個訂單
// @Produce json
// @Param id 			path 	int 	true 	"訂單編號"			minimum(1)
// @Success 200 {object} model.Order 	"成功"
// @Failure 400 {object} errcode.Error 	"請求錯誤"
// @Failure 500 {object} errcode.Error 	"內部錯誤"
// @Router /api/v1/orders/{id} [get]
func (o Order) Get(c *gin.Context) {
	param := service.GetOrderRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	order, err := svc.GetOrder(&param)
	if err != nil {
		global.Logger.Infof("svc.GetOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderFail)
		return
	}

	response.ToResponseAndStatus(order, http.StatusOK)
}

// @Summary 更新訂單
// @Produce json
// @Param id 			path 	int 	true 	"訂單編號"			minimum(1)
// @Param owner 		query 	string 	false	"訂單擁有者" 		maxlength(255)
// @Param type 			query 	int 	false	"訂單種類, 1為買, 2為賣" 		Enums(1, 2)
// @Param quantity 		query 	int 	false	"交易數量"
// @Param price 		query 	int 	false	"交易價格"
// @Param price_policy 	query 	int 	false	"價格原則, 1為限價, 2為市價" 		Enums(1, 2)
// @Success 200 {object} model.Order 	"成功"
// @Failure 400 {object} errcode.Error 	"請求錯誤"
// @Failure 500 {object} errcode.Error 	"內部錯誤"
// @Router /api/v1/orders/{id} [put]
func (o Order) Update(c *gin.Context) {
	getParam := service.GetOrderRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)

	svc := service.New(c.Request.Context())
	order, err := svc.GetOrder(&getParam)
	if err != nil {
		global.Logger.Infof("svc.UpdateOrder on GetOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetOrderFail)
		return
	}

	param := service.UpdateOrderRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	if len(param.Owner) < 3 {
		param.Owner = order.Owner
	}
	if param.Type == 0 {
		param.Type = order.Type
	}
	if param.Quantity == 0 {
		param.Quantity = order.Quantity
	}
	if param.Price == 0 {
		param.Price = order.Price
	}
	if param.PricePolicy == 0 {
		param.PricePolicy = order.PricePolicy
	}

	order, err = svc.UpdateOrder(&param)
	if err != nil {
		global.Logger.Infof("svc.UpdateOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateOrderFail)
		return
	}

	response.ToResponseAndStatus(order, http.StatusOK)
}

// @Summary 刪除訂單
// @Produce json
// @Param id 			path 	int 	true 	"訂單編號"
// @Success 200 {object} model.Order 	"成功"
// @Failure 400 {object} errcode.Error 	"請求錯誤"
// @Failure 500 {object} errcode.Error 	"內部錯誤"
// @Router /api/v1/orders/{id} [delete]
func (o Order) Delete(c *gin.Context) {
	param := service.DeleteOrderRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteOrder(&param)
	if err != nil {
		global.Logger.Infof("svc.DeleteOrder err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteOrderFail)
		return
	}

	response.ToResponse(gin.H{})
}
