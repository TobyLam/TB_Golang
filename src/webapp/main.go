package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!", r.URL.Path)
}

func main() {
	//映射，调用哪个处理器函数
	http.HandleFunc("/", handler)

	//创建路由
	http.ListenAndServe(":8080", nil) // https 调用 http.ListenAndServeTLS()
}
