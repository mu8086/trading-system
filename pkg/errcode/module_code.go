package errcode

var (
	ErrorGetOrderListFail	= NewError(20010001, "取得訂單列表失敗")
	ErrorCreateOrderFail	= NewError(20010002, "建立訂單失敗")
	ErrorGetOrderFail		= NewError(20010003, "取得訂單失敗")
	ErrorUpdateOrderFail	= NewError(20010004, "更新訂單失敗")
	ErrorDeleteOrderFail	= NewError(20010005, "刪除訂單失敗")
	ErrorCountOrderFail		= NewError(20010006, "統計訂單失敗")
)