package model

// 订单项结构
type OrderItem struct {
	OrderItemID int64   //订单项id
	Count       int64   //订单项图书数量
	Amount      float64 //订单项图书金额小计
	Title       string  //订单项图书书名
	Author      string  //订单项图书作者
	Price       float64 //订单项图书价格
	ImgPath     string  //订单项图书封面
	OrderID     string  //所属订单
}
