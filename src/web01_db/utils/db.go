package utils

import (
	"database/sql"
	_ "errors"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:123456@tcp(192.168.1.18:3306)/db2Go")
	if err != nil {
		panic(err.Error())
	}
}
