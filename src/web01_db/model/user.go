package model

import (
	"fmt"
	"web01_db/utils"
)

// User 结构体
type User struct {
	ID     int
	Name   string
	Detail string
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

// GetUserById 根据用户的id从数据库中查询一条记录
func (user *User) GetUserById() (info *User, err error) {
	//sql语句
	sqlStr := "select id,name,detail from go_user where id = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, user.ID)
	//声明
	var id int
	var name string
	var detail string
	err = row.Scan(&id, &name, &detail)
	if err != nil {
		return
	}
	info = &User{
		ID:     id,
		Name:   name,
		Detail: detail,
	}
	return
}

// GetUsers 获取数据库中所有的记录
func (user *User) GetUsers() ([]*User, error) {
	//sql语句
	sqlStr := "select `id`,`name`,`detail` from go_user"
	//执行
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	//创建User切片
	var users []*User
	for rows.Next() {
		//声明
		var id int
		var name string
		var detail string
		err := rows.Scan(&id, &name, &detail)
		if err != nil {
			return nil, err
		}
		u := &User{
			ID:     id,
			Name:   name,
			Detail: detail,
		}
		users = append(users, u)
	}
	return users, err
}
