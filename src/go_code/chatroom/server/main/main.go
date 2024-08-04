package main

import (
	_ "errors"
	"fmt"
	"net"
)

//func readPkg(conn net.Conn) (mes message.Message, err error) {
//	buf := make([]byte, 8096)
//	fmt.Println("读取客户端发送的数据...")
//	//conn.Read 在conn没有被关闭的情况下，才会阻塞
//	//如果客户端关闭了 conn 则不会阻塞
//	_, err = conn.Read(buf[:4])
//	if err != nil {
//		//fmt.Println("conn.Read err=", err)
//		//err = errors.New("read pkg header error")
//		return
//	}
//
//	//fmt.Println("读到的buf=", buf[:4])
//
//	//根据buf[:4] 转换成一个 uint32类型
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[0:4])
//
//	//根据pkgLen读取消息内容
//	n, err := conn.Read(buf[:pkgLen])
//	if n != int(pkgLen) || err != nil {
//		//fmt.Println("conn.Read fail err=", err)
//		//err = errors.New("read pkg body error")
//		return
//	}
//
//	//把pkgLen 反序列化成  -> message.Message
//	//mes需要加&，否则返回的mes为空！！
//	err = json.Unmarshal(buf[:pkgLen], &mes)
//	if err != nil {
//		fmt.Println("json.Unmarshal err=", err)
//		return
//	}
//	return
//
//}

//func writePkg(conn net.Conn, data []byte) (err error) {
//
//	//先发送一个长度给对方
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
//	//发送长度
//	n, err := conn.Write(buf[:4])
//	if n != 4 || err != nil {
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//
//	//发送data本身
//	n, err = conn.Write(data)
//	if n != int(pkgLen) || err != nil {
//		fmt.Println("conn.Write(bytes) fail", err)
//		return
//	}
//	return
//}

//// 编写一个函数serverProcessLogin函数，专门处理登录请求
//func serverProcessLogin(conn net.Conn, mes *message.Message) (err error) {
//	//核心代码...
//	//1.先从mes中取出 mes.Data ,并直接反序列化成LoginMes
//	var loginMes message.LoginMes
//	json.Unmarshal([]byte(mes.Data), &loginMes)
//	if err != nil {
//		fmt.Println("json.Unmarshal fail err=", err)
//		return
//	}
//
//	//1先声明一个 resMes
//	var resMes message.Message
//	resMes.Type = message.LoginResMesType
//
//	//2再声明一个LoginResMes，并完成赋值
//	var loginResMes message.LoginResMes
//
//	//如果用户id=100，密码等于123456，认为合法，否则不合法
//	if loginMes.UserId == 100 && loginMes.UserPwd == "123456" {
//		//合法
//		loginResMes.Code = 200
//	} else {
//		//不合法
//		loginResMes.Code = 500 //500状态码，表示该用户不存在
//		loginResMes.Error = "该用户不存在，请注册再使用..."
//	}
//
//	//3将 loginResMes 序列化
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail", err)
//		return
//	}
//
//	//4.将data赋值给resMes
//	resMes.Data = string(data)
//
//	//5.对resMes进行序列化，准备发送
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("json.Marshal fail", err)
//		return
//	}
//
//	//6.发送data,封装到writePkg函数
//	err = writePkg(conn, data)
//	return
//}

//// 编写一个ServerProcessMes 函数
//// 功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
//func serverProcessMes(conn net.Conn, mes *message.Message) (err error) {
//	switch mes.Type {
//	case message.LoginMesType:
//		//处理登录逻辑
//		err = serverProcessLogin(conn, mes)
//	case message.RegisterMesType:
//	//处理注册
//	default:
//		fmt.Println("消息类型不存在，无法处理...")
//	}
//	return
//}

// 处理和客户端的通讯
func process(conn net.Conn) {

	//需要延时关闭conn
	defer conn.Close()

	//调用总控，创建一个
	processor := &Processor{
		Conn: conn,
	}
	err := processor.process2()
	if err != nil {
		fmt.Println("客户端和服务器端通讯协程错误=err", err)
		return
	}

}

func main() {
	//提示信息
	fmt.Println("服务器[新的结构]在8889端口监听....")
	listen, err := net.Listen("tcp", "0.0.0.0:8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	//一旦监听成功，就等待客户端来连接服务器
	for {
		fmt.Println("等待客户端来连接服务器....")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err=", err)
		}

		//一旦连接成功，则启动一个协程和客户端保持通讯。。。
		go process(conn)

	}
}
