package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
)

// 向数据库中插入订单项
func AddOrderItem(orderItem *model.OrderItem) error {
	//sql语句
	sql := "insert into order_items(count,amount,title,author,price,img_path,order_id) values(?,?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sql, orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID)
	if err != nil {
		return err
	}
	return nil
}
