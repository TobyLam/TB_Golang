package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

// 声明一个结构体，绑定处理用户相关的方法
type UserProcess struct {
}

// 用户注册
func (this *UserProcess) Register(userId int, userPwd string, userName string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务器

	var mes message.Message
	mes.Type = message.RegisterMesType

	//3.创建一个RegisterMes 结构体实例，并赋值
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = userPwd
	registerMes.User.UserName = userName

	//4.将registerMes序列化
	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}

	//7.发送data给服务器端
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册发送信息错误 err=", err)
	}

	//8.接收服务器返回的响应消息
	mes, err = tf.ReadPkg() //mes 就是RegisterResMes
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}

	//将mes的Data部分反序列化成RegisterResMes
	var registerResMes message.RegisterResMes
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请重新登录")
		os.Exit(0)
	} else {
		fmt.Println(registerResMes.Error)
		os.Exit(0)
	}

	return
}

// 用户登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	//1.连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//2.准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType

	//3.创建一个LoginMes 结构体实例，并赋值
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd

	//4.将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}
	//5.把data赋给mes.Data字段
	mes.Data = string(data)

	//6.将mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	//7.发送数据
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送登录消息错误，err=", err)
		return
	}

	//8.接收服务器返回的响应消息
	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserName = loginResMes.UserName
		CurUser.UserStatus = message.UserOnline

		//可以显示当前在线用户列表，遍历 loginResMes.UsersId
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMes.UsersId {

			//如果要求不显示自己在线，增加代码
			if v == userId {
				continue
			}
			fmt.Println("用户id:\t", v)
			//完成 客户端的 onlineUsers 的初始化
			user := &message.User{
				UserId:     v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Print("\n\n")

		//这里还需要在服务器启动一个协程
		//该协程保持和服务器端的通讯，如果服务器有数据推送给客户端
		//则接收并显示在客户端的终端
		go serverProcessMes(conn)

		//1.显示登录成功的菜单..
		for {
			showMenu()
		}

	} else {
		fmt.Println(loginResMes.Error)
	}

	return
}
