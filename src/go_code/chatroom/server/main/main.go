package main

import (
	_ "errors"
	"fmt"
	"go_code/chatroom/server/model"
	"net"
	"time"
)

// 处理和客户端的通讯
func process(conn net.Conn) {

	//需要延时关闭conn
	defer conn.Close()

	//调用总控，创建一个实例
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误 err=", err)
		return
	}

}

func init() {
	//当服务器启动时，就初始化redis连接池
	initPool("192.168.1.18:6379", 16, 0, 300*time.Second)
	initUserDao()
}

// 完成UserDao初始化
func initUserDao() {
	//pool是全局变量
	//注意初始化顺序问题
	//initPool -> initUserDao
	model.MyUserDao = model.NewUserDao(pool)
	model.MySmsDao = model.NewSmsDao(pool)
}

func main() {

	//提示信息
	fmt.Println("服务器正在监听8889端口...")

	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	defer listen.Close() //延时关闭

	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端连接服务器...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯
		go process(conn)

	}
}
