package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
)

// 给关联一个用户登录的方法
type UserProcess struct {
	//字段..

}

// 写一个函数，完成登录
func (this *UserProcess) Login(userId int, userPwd string) (err error) {

	////下一步就要开始定协议...
	//fmt.Printf(" userId = %d userPwd = %s", userId, userPwd)
	//
	//return nil

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
	//3.创建一个LoginMes 结构体
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

	//7.data就是要发送的消息
	//7.1 把data的长度发送给服务器
	//先获取到data的长度->转成一个表示长度的byte切片
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	//发送长度
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	fmt.Printf("客户端，发送消息的长度=%d 内容=%s", len(data), string(data))

	//发送消息本身
	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) fail", err)
		return
	}

	//休眠20
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠了20..")
	//还需要处理服务器端返回的消息
	//创建一个Transfer 实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	mes, err = tf.ReadPkg() //mes 就是
	if err != nil {
		fmt.Println("readPkg(conn) err=", err)
		return
	}
	//将mes的Data部分反序列化成LoginResMes
	var loginResMes message.LoginResMes
	err = json.Unmarshal([]byte(mes.Data), &loginResMes)
	if loginResMes.Code == 200 {
		//fmt.Println("登录成功")

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
