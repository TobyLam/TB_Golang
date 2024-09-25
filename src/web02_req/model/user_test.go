package model

import (
	"fmt"
	"testing"
)

//命令行：
// go test
// go test -v

// TestMain 可以在测试函数执行之前，做一些其他操作
func TestMain(m *testing.M) {
	fmt.Println("测试开始： ")
	//通过m.Run执行来执行测试函数
	m.Run()
}

func TestUser(t *testing.T) {
	fmt.Println("开始测试User中的相关方法")
	//通过t.Run()来执行子测试函数
	//t.Run("测试添加用户：", testAddUser)
	//t.Run("测试获取用户：", testGetUserById)
	t.Run("测试获取用户:", testGetUsers)
}

// 如果函数名不是以Test开头，那么该函数默认不执行,可以设置为一个子测试函数
func testAddUser(t *testing.T) {
	fmt.Println("子测试函数测试添加用户：")
	//user := &User{}

	//调用添加用户的方法
	//user.AddUser()
	//user.AddUser2()
}

// 测试获取一个User
func testGetUserById(t *testing.T) {
	fmt.Println("测试查询一条记录：")
	user := &User{
		ID: 9,
	}
	//调用获取User的方法
	u, _ := user.GetUserById()
	fmt.Println("得到的User的信息时：", u)
}

// 测试获取所有User
func testGetUsers(t *testing.T) {
	fmt.Println("测试查询所有记录：")
	user := &User{
		ID: 9,
	}
	//调用获取所有User的方法
	list, _ := user.GetUsers()

	//遍历
	for k, v := range list {
		fmt.Printf("第%d个用户是%v\n", k+1, v)
	}
}
