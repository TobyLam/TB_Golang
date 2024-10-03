package dao

import (
	"bookstore001/model"
	"bookstore001/utils"
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
