package main

import (
	"fmt"
	"net/http"
)

// 创建处理器函数
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "你发送的请求的请求地址是：", r.URL.Path)
	fmt.Fprintln(w, "你发送的请求的请求地址后的查询字符串是：", r.URL.RawQuery)
	fmt.Fprintln(w, "请求头中的所有信息", r.Header)
	fmt.Fprintln(w, "请求头中Accept-Encoding的信息有：", r.Header["Accept-Encoding"])
	fmt.Fprintln(w, "请求头中Accept-Encoding的属性值是：", r.Header.Get("Accept-Encoding"))

	//获取请求体的内容有两种方式，只能选其一

	//获取请求体的内容，方式一【body】：
	/**
	//获取请求体中的内容长度
	len := r.ContentLength
	//创建byte切片
	body := make([]byte, len)
	//将请求体中的内容读到body中
	r.Body.Read(body)
	//在浏览器中显示请求体中的内容
	fmt.Fprintln(w, "请求体中的内容有：", string(body))
	**/

	// 方式二[form]：

	//解析表单.在调用r.Form之前，必须执行该操作
	//表单编码是multipart/form-data（文件上传）时，使用MultipartForm字段
	//r.ParseForm()
	//r.ParseMultipartForm(20)
	//获取请求参数
	//如果form表单的action属性的URL地址中也有与form表单参数名相同的请求参数，
	//参数都可以得到，且form表单中的参数值在URL的参数值前面
	//fmt.Fprintln(w, "请求参数有：", r.Form)
	//fmt.Fprintln(w, "POST请求的form表单中的请求参数有：", r.PostForm)
	//fmt.Fprintln(w, "multipart/form-data编码请求参数有:", r.MultipartForm)

	//通过直接调用FormValue方法和PostFormValue方法直接获取请求参数的值
	fmt.Fprintln(w, "URL中的user请求参数的值是：", r.FormValue("user"))
	fmt.Fprintln(w, "URL中的username请求参数的值是：", r.PostFormValue("username"))

}

func main() {

	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
