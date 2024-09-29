package main

import (
	"fmt"
	"net/http"
)

// 设置cookie
func setCookie(w http.ResponseWriter, r *http.Request) {
	//创建Cookie
	cookie := http.Cookie{
		Name:     "user",
		Value:    "admin1",
		HttpOnly: true,
		MaxAge:   60,
	}
	cookie2 := http.Cookie{
		Name:     "userw",
		Value:    "adminw",
		HttpOnly: true,
	}
	//将Cookie发送给浏览器
	//w.Header().Set("Set-Cookie", cookie.String())
	//添加第二个Cookie
	//w.Header().Add("Set-Cookie", cookie2.String())
	//直接调用http的SetCookie函数设置Cookie
	http.SetCookie(w, &cookie)
	http.SetCookie(w, &cookie2)
}

// 获取cookie
func getCookies(w http.ResponseWriter, r *http.Request) {
	//获取请求头中所有的Cookie
	//cookies := r.Header["Cookie"]

	//获取某一个cookie
	cookie, _ := r.Cookie("user")
	fmt.Println("得到的Cookie有：", cookie)
}

func main() {

	http.HandleFunc("/setCookie", setCookie)
	http.HandleFunc("/getCookies", getCookies)

	http.ListenAndServe(":8080", nil)
}
