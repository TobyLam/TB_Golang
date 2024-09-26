package main

import (
	"bookstore001/controller"
	"html/template"
	"net/http"
)

// IndexHandler 去首页
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	//解析模板
	t := template.Must(template.ParseFiles("views/index.html"))
	//执行
	t.Execute(w, "")
}

func main() {
	//设置处理静态资源,如css和js
	// http.Dir()使用绝对路径views/static，而非/views/static
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))

	http.HandleFunc("/main", IndexHandler)

	//去登陆
	http.HandleFunc("/login", controller.Login)
	//去注册
	http.HandleFunc("/regist", controller.Regist)

	http.ListenAndServe(":8080", nil)
}
