package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
)

// 向购物项表中插入购物项
func AddCartItem(cartItem *model.CartItem) error {
	//sql语句
	sqlStr := "insert into cart_items(count,amount,book_id,cart_id) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID)
	if err != nil {
		return err
	}
	return nil
}

// 根据图书id获取对应的购物项
func GetCartItemByBookIDAndCartID(bookID string, cartID string) (*model.CartItem, error) {
	//sql语句
	sqlStr := "select id,count,amount,cart_id from cart_items where book_id = ? and cart_id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, bookID, cartID)
	//创建cartItem
	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	if err != nil {
		return nil, err
	}
	return cartItem, nil
}

// 根据购物车id获取购物车所有购物项
func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	//sql查询
	sqlStr := "select id,count,amount,book_id,cart_id from cart_items where cart_id = ?"
	//执行
	rows, err := utils.Db.Query(sqlStr, cartID)
	if err != nil {
		return nil, err
	}
	var cartItems []*model.CartItem
	for rows.Next() {
		//设置一个变量接收bookId
		var bookID string
		//创建cartItem
		cartItem := &model.CartItem{}
		err2 := rows.Scan(&cartItem.CartItemID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID)
		if err2 != nil {
			return nil, err2
		}
		//根据bookID获取图书信息
		book, _ := GetBookByID(bookID)
		//将book设置到购物项中
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	return cartItems, nil
}
