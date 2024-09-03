package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// 数据库连接配置
const (
	DBHost     = "192.168.1.18"
	DBPort     = 3306
	DBUser     = "root"
	DBPassword = "123456"
	DBName     = "db2Go"
)

func main() {

	//构建DSN
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", DBUser, DBPassword, DBHost, DBPort, DBName)

	//创建数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	//延时关闭
	defer db.Close()

	//检查是否连接成功
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Successfully connected to database.")
	}

	//使用连接池
	var maxOpenConns = 10 //最大连接数
	var maxIdleConns = 5  //最大限制连接数

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)

	//新增一条数据
	insertStmt := `INSERT INTO go_user (name) VALUES (?)`
	result, err := db.Exec(insertStmt, "烧糊的狐小苏")
	if err != nil {
		log.Fatal(err)
	}

	//获取新增的id
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Insert new user with ID:%d\n", lastInsertID)

	//处理事务
	tx, err := db.Begin()
	if err != nil {
		return
	}

	//在事务中执行操作
	_, err = tx.Exec("UPDATE go_user SET name=? WHERE id=?", "黄子弘凡", 4)
	if err != nil {
		tx.Rollback()
		return
	}

	//提交事务
	err = tx.Commit()
	if err != nil {
		return
	}

	//更新操作
	updateStmt := `UPDATE go_user SET name=? WHERE id=?`
	res, err := db.Exec(updateStmt, "福山", lastInsertID)
	if err != nil {
		log.Fatal(err)
	}
	affectCount, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d rows affected\n", affectCount)

}
