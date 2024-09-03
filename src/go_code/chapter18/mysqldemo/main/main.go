package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	//DSN (Data Source Name) 数据库连接信息字符串
	dsn := "root:123456@tcp(192.168.1.18:3306)/db2Go?charset=utf8mb4&parseTime=True&loc=Local"

	// 创建数据连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	//延时关闭连接
	defer db.Close()

	//检查连接是否有效
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("connected to the database")
	}

	//执行查询
	//rows, err := db.Query("select * from go_user where id=?", 1)
	rows, err := db.Query("select * from go_user")

	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//循环遍历
	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}

	//插入或更新操作
	_, err = db.Exec("Insert into go_user (name) values (?)", "诸葛正我")
	if err != nil {
		log.Fatal(err)
	}
}
