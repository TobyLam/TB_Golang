package model

// 购物项结构体
type CartItem struct {
	CartItemID int64   //购物项的id
	Book       *Book   //购物项图书信息
	Count      int64   //图书数量
	Amount     float64 //图书金额小计，可计算获得
	CartID     string  //当前购物项所属购物车
}

// 获取购物项中图书的金额小计，由价格和数量计算得到
func (this *CartItem) GetAmount() float64 {
	//获取当前购物项中图书的价格
	price := this.Book.Price
	return float64(this.Count) * price
}
