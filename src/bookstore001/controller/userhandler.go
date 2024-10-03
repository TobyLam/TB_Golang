package controller

import (
	"bookstore001/dao"
	"bookstore001/model"
	"bookstore001/utils"
	"html/template"
	"net/http"
)

// Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名、密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")

	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserNameAndPassword(username, password)
	if user.Id > 0 {
		//用户名和密码正确
		//生成UUID作为Session的id
		uuid := utils.CreateUUID()
		//创建一个Session
		session := &model.Session{
			SessionID: uuid,
			UserName:  user.Username,
			UserID:    user.Id,
		}
		//将Session保存到数据库中
		dao.AddSession(session)
		//创建一个Cookie，与Session关联
		cookie := http.Cookie{
			Name:     "user",
			Value:    uuid,
			HttpOnly: true,
		}
		//将Cookie发送给浏览器
		http.SetCookie(w, &cookie)

		t := template.Must(template.ParseFiles("views/pages/user/login_success.html"))
		t.Execute(w, user)
	} else {
		//用户名或密码不正确
		t := template.Must(template.ParseFiles("views/pages/user/login.html"))
		t.Execute(w, "用户名或密码不正确！")
	}
}

// 处理用户注销的函数
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中对应的session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	GetPageBooksByPrice(w, r)
}

// Regist 处理用户注册的函数
func Regist(w http.ResponseWriter, r *http.Request) {
	//获取用户名、密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")

	//调用userdao中验证用户名的方法
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，将用户保存到数据库中
		err := dao.SaveUser(username, password, email)
		if err != nil {
			t := template.Must(template.ParseFiles("views/pages/user/regist.html"))
			t.Execute(w, "")
		} else {
			t := template.Must(template.ParseFiles("views/pages/user/regist_success.html"))
			t.Execute(w, "")
		}

	}
}

// CheckUserName 通过发送Ajax验证用户名是否可用
func CheckUserName(w http.ResponseWriter, r *http.Request) {
	//获取用户输入的用户名
	username := r.PostFormValue("username")
	//调用userdao中验证用户名的方法
	user, _ := dao.CheckUserName(username)
	if user.Id > 0 {
		//用户名已存在
		w.Write([]byte("用户名已存在！"))
	} else {
		//用户名可用
		w.Write([]byte("<font style='color:green'>用户名可用！</font>"))
	}
}
