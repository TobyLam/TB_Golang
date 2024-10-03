package dao

import (
	"bookstore001/model"
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("测试bookdao中的方法")
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
	fmt.Println("测试bookdao中的相关函数")
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
	fmt.Println("测试Session相关函数")
	//t.Run("测试添加Session", testAddSession)
	//t.Run("测试删除Session", testDeleteSession)
	t.Run("测试获取Session", testGetSession)
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
