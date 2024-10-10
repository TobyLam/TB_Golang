package dao

import (
	"bookstore001/model"
	"fmt"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	//fmt.Println("测试bookdao中的方法")
	m.Run()
}

func TestUser(t *testing.T) {
	//fmt.Println("测试userdao中的函数")
	//t.Run("验证用户名或密码：", testLogin)
	//t.Run("验证用户名", testRegist)
	//t.Run("保存用户：", testSave)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("钟离", "123456")
	fmt.Println("获取用户信息是：", user)
}

func testRegist(t *testing.T) {
	user, _ := CheckUserName("钟离")
	fmt.Println("获取用户信息是：", user)
}

func testSave(t *testing.T) {
	SaveUser("钟离", "123456", "yanwangdijun@liyue.com")
}

func TestBook(t *testing.T) {
	//fmt.Println("测试bookdao中的相关函数")
	//t.Run("测试获取所有图书", testGetBooks)
	//t.Run("测试添加图书", testAddBook)
	//t.Run("测试删除图书", testDeleteBook)
	//t.Run("测试获取一本图书", testGetBook)
	//t.Run("测试更新图书", TestUpdateBook)
	//t.Run("测试获取分页的图书", testGetPageBooks)
	//t.Run("测试带分页和价格范围的图书", testGetPageBooksByPrice)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%d本图书的信息是：%v\n", k+1, v)
	}
}

func testAddBook(t *testing.T) {
	b := &model.Book{
		Title:   "封神演绎",
		Author:  "许仲琳",
		Price:   100,
		Sales:   1000,
		Stock:   2000,
		ImgPath: "/static/img/default.jpg",
	}
	//调用添加图书的函数
	AddBook(b)
}

func testDeleteBook(t *testing.T) {
	//调用删除图书的函数
	DeleteBook("47")
}

func testGetBook(t *testing.T) {
	//调用获取图书的函数
	book, _ := GetBookByID("45")
	fmt.Println("获取到的图书信息:", book)
}

func TestUpdateBook(t *testing.T) {
	b := &model.Book{
		ID:      45,
		Title:   "封神演绎",
		Author:  "明·许仲琳",
		Price:   100,
		Sales:   1000,
		Stock:   2000,
		ImgPath: "/static/img/default.jpg",
	}
	//调用修改图书的函数
	UpdateBook(b)
}

func testGetPageBooks(t *testing.T) {
	page, _ := GetPageBooks("9")
	fmt.Println(" 当前页是:", page.PageNo)
	fmt.Println(" 总页数是:", page.TotalPageNo)
	fmt.Println(" 总记录数是:", page.TotalRecord)
	fmt.Println(" 当前页中的图书有:")
	for i, v := range page.Books {
		fmt.Printf("第%d本图书信息是：%v\n", i+1, v)
	}
}

func testGetPageBooksByPrice(t *testing.T) {
	page, _ := GetPageBooksByPrice("1", "10", "30")
	fmt.Println(" 当前页是:", page.PageNo)
	fmt.Println(" 总页数是:", page.TotalPageNo)
	fmt.Println(" 总记录数是:", page.TotalRecord)
	fmt.Println(" 当前页中的图书有:")
	for i, v := range page.Books {
		fmt.Printf("第%d本图书信息是：%v\n", i+1, v)
	}
}

func TestSession(t *testing.T) {
	//fmt.Println("测试Session相关函数")
	//t.Run("测试添加Session", testAddSession)
	//t.Run("测试删除Session", testDeleteSession)
	//t.Run("测试获取Session", testGetSession)
}

func testAddSession(t *testing.T) {
	sess := &model.Session{
		SessionID: "125648752365421",
		UserName:  "宋江",
		UserID:    9,
	}
	AddSession(sess)
}

func testDeleteSession(t *testing.T) {
	DeleteSession("125648752365421")
}

func testGetSession(t *testing.T) {
	sess, _ := GetSession("4a460c09-c69c-40ef-6c65-c728acd4b959")
	fmt.Println("Session的信息是：", sess)
}

func TestCart(t *testing.T) {
	//fmt.Println("测试购物车的相关函数")
	//t.Run("测试添加购物车", testAddCart)
	//t.Run("测试根据图书的id获取对应的购物项", testGetCartItemByBookID)
	//t.Run("测试根据购物车的id获取所有的购物项", testGetItemsByCartID)
	//t.Run("测试根据用户id获取对应的购物车", testGetCartByUserID)
	//t.Run("测试根据图书id和购物车id更新购物车、购物项", testUpdateBookCount)
	//t.Run("根据购物车的id删除购物车", testDeleteCartByCartID)
	//t.Run("测试根据购物项id删除购物项", testDeleteCartItemByID)
}

func testAddCart(t *testing.T) {
	//设置要购买的第一本书
	book := &model.Book{
		ID:    1,
		Price: 27.20,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 23.00,
	}
	//创建购物项切片
	var cartItems []*model.CartItem
	//创建两个购物项
	cartItem := &model.CartItem{

		Book:   book,
		Count:  10,
		CartID: "6668888",
	}
	cartItems = append(cartItems, cartItem)
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  10,
		CartID: "6668888",
	}
	cartItems = append(cartItems, cartItem2)
	//创建购物车
	cart := &model.Cart{
		CartID:    "6668888",
		CartItems: cartItems,
		UserID:    17,
	}
	AddCart(cart)
}

func testGetCartItemByBookID(t *testing.T) {
	cartItem, _ := GetCartItemByBookIDAndCartID("1", "6668888")
	fmt.Println("图书id=1的购物项的信息是：", cartItem)
}

func testGetItemsByCartID(t *testing.T) {
	cartItems, _ := GetCartItemsByCartID("6668888")
	for k, v := range cartItems {
		fmt.Printf("第%v个购物项是：%v\n", k+1, v)
	}
}

func testGetCartByUserID(t *testing.T) {
	cart, _ := GetCartByUserID(17)
	fmt.Println("id为17的用户的购物车信息是：", cart)
}

func testUpdateBookCount(t *testing.T) {
	//UpdateBookCount(10, 1, "f2ee85a3-bb35-4fe7-7d84-6cc55726eaef")
}

func testDeleteCartByCartID(t *testing.T) {
	DeleteCartByCartID("57b00c70-8040-41d0-7dde-8fb31d85a6b0")
}

func testDeleteCartItemByID(t *testing.T) {
	DeleteCartItemByID("1")
}

func TestOrder(t *testing.T) {
	fmt.Println("测试订单相关函数")
	//t.Run("测试添加订单和订单项", testAddOrder)
	//t.Run("测试获取所有订单", testGetOrders)
	//t.Run("测试获取所有的订单项", testGetOrderItems)
	t.Run("测试根据用户id获取订单", testGetMYOrders)
}

func testAddOrder(t *testing.T) {
	//生成订单号
	orderID := "18826418290"
	//下单时间
	timeStr := time.Now().Format("2006-01-02 15:04:05")
	//创建订单
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  timeStr,
		TotalCount:  2,
		TotalAmount: 400,
		State:       0,
		UserID:      18,
	}
	//创建订单项
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  300,
		Title:   "封神演义",
		Author:  "许仲琳",
		Price:   300,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  100,
		Title:   "红楼梦",
		Author:  "曹雪芹/高鹗",
		Price:   100,
		ImgPath: "/static/img/default.jpg",
		OrderID: orderID,
	}
	//保存订单
	AddOrder(order)
	//保存订单项
	AddOrderItem(orderItem)
	AddOrderItem(orderItem2)
}

func testGetOrders(t *testing.T) {
	orders, _ := GetOrders()
	for _, v := range orders {
		fmt.Println("订单信息是:", v)
	}
}

func testGetOrderItems(t *testing.T) {
	ordersItems, _ := GetOrderItemsByOrderID("69fd225c-dda8-44ff-443b-10c9063ee69c")
	for _, v := range ordersItems {
		fmt.Println("订单项信息是:", v)
	}
}

func testGetMYOrders(t *testing.T) {
	orders, _ := GetMyOrders(18)
	for _, v := range orders {
		fmt.Println("订单信息是:", v)
	}
}
