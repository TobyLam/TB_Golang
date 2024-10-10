package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
)

// 向数据库中插入订单
func AddOrder(order *model.Order) error {
	//sql语句
	sql := "insert into orders(id,create_time,total_count,total_amount,state,user_id) values(?,?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sql, order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.State, order.UserID)
	if err != nil {
		return err
	}
	return nil
}
