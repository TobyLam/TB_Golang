package process2

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/utils"
	"net"
)

type SmsProcess struct {
	//字段..
}

/**
 * 发送消息
 */
func (this *SmsProcess) SendMes(mes *message.Message) (err error) {
	//取出mes的内容赋值给SmsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	//根据接收者id区分私聊和群发
	if smsMes.ReceiverId == 0 {
		err = this.SendGroupMes(mes)
	} else {
		err = this.SendPrivateMes(mes)
	}
	return
}

/**
 * 转发私聊消息
 */
func (this *SmsProcess) SendPrivateMes(mes *message.Message) (err error) {

	//取出mes的内容赋值给SmsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//根据用户id，从在线用户列表中获取链接Conn
	up, ok := userMgr.onlineUsers[smsMes.ReceiverId]
	if !ok {
		err = errors.New("receiverId conn not exist")
		return
	}

	//给在线用户发送消息
	err = this.SendMesToEachOnlineUser(&smsMes, up.Conn)
	if err != nil {
		fmt.Printf("私聊消息给用户(id:%d)失败\n", smsMes.ReceiverId)
	}

	return
}

/*
 * 群体转发消息
 * 排除发送者
 */
func (this *SmsProcess) SendGroupMes(mes *message.Message) (err error) {

	//取出mes的内容赋值给SmsMes
	var smsMes message.SmsMes
	err = json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	//遍历服务器端的onlineUsers map[int]*UserProcess
	//将消息转发出去
	for id, up := range userMgr.onlineUsers {
		//过滤掉发送者
		if id == smsMes.UserId {
			continue
		}
		//给在线用户发送消息
		err = this.SendMesToEachOnlineUser(&smsMes, up.Conn)
		if err != nil {
			fmt.Printf("群发消息给用户(id:%d)失败\n", id)
		}
	}
	return
}

/**
 * 给在线用户发送消息
 */
func (this *SmsProcess) SendMesToEachOnlineUser(smsMes *message.SmsMes, conn net.Conn) (err error) {

	//创建一个发送消息响应实例
	smsResMes := &message.SmsResMes{
		Content:    smsMes.Content,
		User:       smsMes.User,
		ReceiverId: smsMes.ReceiverId,
	}

	//序列化smsResMes
	data, err := json.Marshal(smsResMes)
	if err != nil {
		fmt.Println("json.Marshal(smsResMes) err=", err)
		return
	}

	//创建一个message实例，返回消息类型为发送消息响应的数据
	var mes message.Message
	mes.Type = message.SmsResMesType
	mes.Data = string(data)

	//序列化mes
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal() err=", err)
		return

	}

	//创建一个Transfer 实例 ，发送data
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("服务器转发消息失败 err=", err)
	}

	return
}
