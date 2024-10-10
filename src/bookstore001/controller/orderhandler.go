package controller

import (
	"bookstore001/dao"
	"bookstore001/model"
	"bookstore001/utils"
	"html/template"
	"net/http"
	"time"
)

// 去结账
func Checkout(w http.ResponseWriter, r *http.Request) {
	//获取session
	_, session := dao.IsLogin(r)
	//获取用户的id
	userID := session.UserID
	//获取购物车
	cart, _ := dao.GetCartByUserID(userID)
	//生成订单号
	orderID := utils.CreateUUID()
	//下单时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//创建Order
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		State:       0,
		UserID:      int64(userID),
	}
	//将订单保存到数据库
	dao.AddOrder(order)
	//保存订单项
	//获取购物车中的购物项
	cartItems := cart.CartItems
	//遍历每一个购物项
	for _, v := range cartItems {
		//创建orderItem
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}
		//保存购物项到数据库
		dao.AddOrderItem(orderItem)
		//更新当前购物项中图书的库存和销量
		book := v.Book
		book.Sales = book.Sales + int(v.Count)
		book.Stock = book.Stock - int(v.Count)
		//更新图书信息
		dao.UpdateBook(book)
	}
	//清空购物车
	dao.DeleteCartByCartID(cart.CartID)
	//将订单号设置到session中
	session.OrderID = orderID
	//解析模板
	t := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	//执行
	t.Execute(w, session)
}
