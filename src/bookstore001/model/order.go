package model

// 订单结构体
type Order struct {
	OrderID     string  //订单号
	CreateTime  string  //生成订单时间
	TotalCount  int64   //订单图书总数
	TotalAmount float64 //订单图书总金额
	State       int64   //订单状态 0 未发货 1 已发货 2 交易完成
	UserID      int64   //订单所属的用户
}
