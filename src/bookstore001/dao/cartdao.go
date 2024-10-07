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

// 根据用户id获取购物车
func GetCartByUserID(userID int) (*model.Cart, error) {
	//sql语句
	sql := "select id,total_count,total_amount,user_id from carts where user_id = ?"
	//执行
	row := utils.Db.QueryRow(sql, userID)
	//创建购物车实例
	cart := &model.Cart{}
	//扫描赋值
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	if err != nil {
		return nil, err
	}

	//获取当前购物车里面所有购物项
	cartItems, _ := GetCartItemsByCartID(cart.CartID)
	//将所有购物项设置到购物车中
	cart.CartItems = cartItems
	return cart, nil
}

// 更新购物车中图书的总数量、总金额
func UpdateCart(cart *model.Cart) error {
	//sql语句
	sql := "update carts set total_count = ? ,total_amount = ? where id = ?"
	//执行
	_, err := utils.Db.Exec(sql, cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID)
	if err != nil {
		return err
	}
	return nil
}
