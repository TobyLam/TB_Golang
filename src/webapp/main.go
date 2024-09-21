package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "Hello World!", r.URL.Path)
	fmt.Fprintln(w, "通过自己创建的多路复用器处理请求", r.URL.Path)
}

func main() {
	//创建多路复用器
	mux := http.NewServeMux()

	//映射，调用哪个处理器函数
	//http.HandleFunc("/", handler)
	mux.HandleFunc("/", handler)

	//创建路由
	//http.ListenAndServe(":8080", nil) // https 调用 http.ListenAndServeTLS()
	http.ListenAndServe(":8080", mux) // https 调用 http.ListenAndServeTLS()

}
