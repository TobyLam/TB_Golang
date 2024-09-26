package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
)

// 根据用户名和密码从数据库钟查询一条记录
func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	return user, nil
}

// 根据用户名从数据库中查询一条记录
func CheckUserName(username string) (*model.User, error) {
	//sql语句
	sqlStr := "select id,username,password,email from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	//Scan（）字段必须对应上sqlStr中查询的所有字段
	row.Scan(&user.Id, &user.Username, &user.Password, &user.Email)

	return user, nil
}

// 向数据库中插入用户信息
func SaveUser(username, password, email string) error {
	//sql语句
	sqlStr := "insert into users(username,password,email) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
