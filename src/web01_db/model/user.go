package model

import (
	"fmt"
	"web01_db/utils"
)

// User 结构体
type User struct {
	ID   int
	Name string
}

// AddUser 添加User的方法一
func (user *User) AddUser() (err error) {
	//1.sql语句
	sqlStr := "insert into go_user(name,detail) values (?,?)"
	//2.预编译
	instmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译出现异常:", err)
		return
	}
	//3.执行
	_, err = instmt.Exec("钟离", "岩王帝君")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return
	}
	return
}

// AddUser2 添加User的方法二
func (user *User) AddUser2() (err error) {
	//1.sql语句
	sqlStr := "insert into go_user(name,detail) values (?,?)"

	//2.执行
	_, err = utils.Db.Exec(sqlStr, "枫原万叶", "可叹！落叶飘零~")
	if err != nil {
		fmt.Println("执行出现异常:", err)
		return
	}
	return
}
