package main

import (
	_ "encoding/json"
	_ "fmt"
	"html/template"
	"net/http"
	_ "web02_req/model"
)

// 创建处理器函数

func testTemplate(w http.ResponseWriter, r *http.Request) {
	//解析模板文件
	//t, _ := template.ParseFiles("index.html")
	//通过Must函数让Go自动处理异常
	t := template.Must(template.ParseFiles("index.html", "index2.html"))
	//执行
	//t.Execute(w, "Hello Template")
	//将响应数据在index2.html文件中显示
	t.ExecuteTemplate(w, "index2.html", "原神启动")
}

func main() {

	http.HandleFunc("/testTemplate", testTemplate)

	http.ListenAndServe(":8080", nil)
}
