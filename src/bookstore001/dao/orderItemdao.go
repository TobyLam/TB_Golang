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

// 根据订单号获取该订单所有订单项
func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	//sql语句
	sql := "select id,count,amount,title,author,price,img_path,order_id from order_items where order_id = ?"
	//执行
	rows, err := utils.Db.Query(sql, orderID)
	if err != nil {
		return nil, err
	}
	//创建切片
	var orderItems []*model.OrderItem
	//遍历每一个订单项
	for rows.Next() {
		orderItem := &model.OrderItem{}
		//扫描赋值
		rows.Scan(&orderItem.OrderItemID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID)
		//添加到切片中
		orderItems = append(orderItems, orderItem)
	}

	return orderItems, nil
}
