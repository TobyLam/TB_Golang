package dao

import (
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
	t.Run("测试获取所有图书", testGetBooks)
}

func testGetBooks(t *testing.T) {
	books, _ := GetBooks()
	//遍历得到每一本图书
	for k, v := range books {
		fmt.Printf("第%d本图书的信息是：%v\n", k+1, v)
	}
}
