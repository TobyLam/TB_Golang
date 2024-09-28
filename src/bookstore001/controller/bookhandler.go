package controller

import (
	"bookstore001/dao"
	"html/template"
	"net/http"
)

// 获取所有图书
func GetBooks(w http.ResponseWriter, r *http.Request) {
	//调用bookdao获取所有图书的函数
	books, _ := dao.GetBooks()
	//解析模板文件
	t := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	//执行
	t.Execute(w, books)
}
