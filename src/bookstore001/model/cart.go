package model

// 购物车结构体
type Cart struct {
	CartID      string      //购物车id
	CartItems   []*CartItem //购物车中所有的购物项
	TotalCount  int64       //购物车中图书的总数量，可计算获得
	TotalAmount float64     //购物车中图书的总金额，可计算获得
	UserID      int         //购物车所属的用户
}

// 获取购物车图书的总数量
func (this *Cart) GetTotalCount() int64 {
	var totalCount int64
	//遍历购物车中的购物项切片
	for _, v := range this.CartItems {
		totalCount += v.Count
	}
	return totalCount
}

// 获取购物车图书的总金额
func (this *Cart) GetTotalAmount() float64 {
	var totalAmount float64
	//遍历购物车中的购物项切片
	for _, v := range this.CartItems {
		totalAmount += totalAmount + v.Amount
	}
	return totalAmount
}
