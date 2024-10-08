package controller

import (
	"bookstore001/dao"
	"bookstore001/model"
	"bookstore001/utils"
	_ "fmt"
	"net/http"
)

// 添加图书到购物车
func AddBook2Cart(w http.ResponseWriter, r *http.Request) {
	//判断是否登录
	flag, session := dao.IsLogin(r)
	if flag {
		//已经登录

		//获取要添加的图书id
		bookID := r.FormValue("bookId")
		//根据图书id获取图书信息
		book, _ := dao.GetBookByID(bookID)

		//获取用户的id
		userID := session.UserID
		//判断数据库中是否有当前用户的购物车
		cart, _ := dao.GetCartByUserID(userID)
		if cart != nil {
			//当前用户已经有购物车，需要判断购物车中是否有当前这本图书
			carItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)
			if carItem != nil {
				//购物车的购物项中已经有该图书,只需要修改该购物项所对应的数量、价格
				//1.获取购物车切片中的所有购物项
				cts := cart.CartItems
				//2.遍历每一个购物项
				for _, v := range cts {
					//3.找到当前的购物项
					if v.Book.ID == carItem.Book.ID {
						//将购物项中的图书的数量加1
						v.Count = v.Count + 1
						//更新数据库中该购物项的图书数量
						dao.UpdateBookCount(v.Count, v.Book.ID, cart.CartID)
					}
				}
			} else {
				//购物车的购物项中还没有该图书，需要创建一个购物项并添加到数据库中
				//创建购物车中的购物项
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}
				//将购物项添加到当前cart的切片中
				cart.CartItems = append(cart.CartItems, cartItem)
				//将新创建的购物项添加到数据库中
				dao.AddCartItem(cartItem)
			}
			//无论之前购物车中是否有当前图书购物项，都需要更新购物车中的总数量、总金额
			dao.UpdateCart(cart)
		} else {
			//当前用户还没有购物车，需要创建一个购物车并添加到数据库中

			//1.创建购物车
			//生成购物车的id
			cartID := utils.CreateUUID()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}

			//2.创建购物车中的购物项
			//声明一个CartItem类型的切片
			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				Amount: 1,
				CartID: cartID,
			}
			//购物项添加到切片中
			cartItems = append(cartItems, cartItem)

			//3.切片设置到cart中
			cart.CartItems = cartItems

			//4.将购物车cart保存到数据库中
			dao.AddCart(cart)
		}
		w.Write([]byte("您刚刚将" + book.Title + "添加到了购物车！"))
	} else {
		//没有登录
		w.Write([]byte("请先登录！"))
	}

}
