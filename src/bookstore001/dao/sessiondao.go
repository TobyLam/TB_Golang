package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
	"net/http"
)

// 添加Session
func AddSession(sess *model.Session) error {
	//sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

// 删除Session
func DeleteSession(sessID string) error {
	//sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

// 获取Session
func GetSession(sessID string) (*model.Session, error) {
	//sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建Session
	sess := &model.Session{}
	//扫描赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}

// 判断用户是否已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据cookie的name获取cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库查询对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}
