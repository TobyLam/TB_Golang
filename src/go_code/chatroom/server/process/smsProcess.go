package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
	//字段..
}

/*
 * 群体转发消息
 * 排除发送者
 */
func (this *SmsProcess) SendGroupMes(mes *message.Message) {

	//取出mes的内容赋值给SmsMes
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//序列化
	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
	}

	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发出去
	for id, up := range userMgr.onlineUsers {
		//过滤掉发送者
		if id == smsMes.UserId {
			continue
		}
		//给在线用户发送消息
		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

/**
 * 给在线用户发送消息
 */
func (this *SmsProcess) SendMesToEachOnlineUser(data []byte, conn net.Conn) {
	//创建一个Transfer 实例 ，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("服务器转发消息失败 err=", err)
	}
}
