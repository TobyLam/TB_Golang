package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
)

// 向购物车表中插入购物车
func AddCart(cart *model.Cart) error {
	//sql语句
	sqlStr := "insert into carts(id,total_count,total_amount,user_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID)
	if err != nil {
		return err
	}
	//获取购物车中的所有购物项
	cartItems := cart.CartItems
	//遍历得到每个购物项
	for _, cartItem := range cartItems {
		//将购物项插入到数据库中
		AddCartItem(cartItem)
	}
	return nil
}
