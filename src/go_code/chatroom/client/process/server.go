package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
	"net"
	"os"
)

// 显示登录成功后的界面..
func showMenu() {
	fmt.Println("-----------请选择----------")
	fmt.Println("--------1. 显示在线用户列表----------")
	fmt.Println("--------2. 私聊----------")
	fmt.Println("--------3. 群发----------")
	fmt.Println("--------4. 信息列表----------")
	fmt.Println("--------5. 退出系统----------")
	fmt.Println("请选择（1-5）：")

	var key int
	var content string //发送消息内容
	var receiverId int //接收者id
	//因为总会使用到SmsProcess实例，因此定义在switch外部
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:

		//fmt.Println("显示在线用户列表-")
		outputOnlineUser()

	case 2:
		// 显示在线用户列表：
		outputOnlineUser()
		// 接收私聊用户id
		fmt.Println("请输入私聊用户id：")
		fmt.Scanf("%d\n", &receiverId)
		// 接收消息内容
		fmt.Println("请输入需要发送的内容:)")
		fmt.Scanf("%s\n", &content)
		// 发送私聊信息
		smsProcess.SendPrivateMes(content, receiverId)

	case 3:
		// 接收群发消息内容
		fmt.Println("说点什么吧：）")
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMes(content)

	case 4:
		fmt.Println("信息列表")

	case 5:
		fmt.Println("你选择退出了系统...")

		os.Exit(0)
	default:
		fmt.Println("你输入的选项不正确...")
	}

}

// 和服务器保持通讯
func serverProcessMes(conn net.Conn) {

	// 1.创建一个transfer实例,读取服务器发送的信息
	tf := &utils.Transfer{
		Conn: conn,
	}

	// 2.死循环，读取服务器发送的消息
	for {
		//fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err=", err)
			return
		}
		//如果读取到消息，下一步处理逻辑
		switch mes.Type {
		case message.NotifyUserStatusMesType: //有人上线了

			//1、取出NotifyUserStatusMes
			var notifyUserStatusMes message.NotifyUserStatusMes
			err = json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
			//2. 把这个用户的信息，保存到客户端的map[int]User中
			updateUserStatus(&notifyUserStatusMes)

		case message.SmsResMesType: //有人发送消息

			//显示消息
			err = outputMes(&mes)

		default:
			fmt.Println("服务器端返回了未知消息类型")
		}
	}
}
