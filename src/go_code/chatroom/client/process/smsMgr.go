package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
)

/**
 * 根据实际情况
 * 分别渲染私聊消息和群聊消息
 */
func outputMes(mes *message.Message) (err error) {
	//反序列化mes.Data
	var smsResMes message.SmsResMes
	err = json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.Unmarshal(smsResMes) err", err.Error())
		return
	}
	//判断响应消息中是否指定接收用户
	fmt.Println(smsResMes.ReceiverId)
	if smsResMes.ReceiverId == 0 {
		outputGroupMes(mes)
	} else {
		outputPrivateMes(mes)
	}

	return
}

/**
 * 显示私聊信息
 */
func outputPrivateMes(mes *message.Message) { //mes 一定是 smsResMes类型
	//显示即可
	//1.反序列化mes.Data
	var smsResMes message.SmsResMes
	err := json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("【私聊】%s:\t%s", smsResMes.UserName, smsResMes.Content)
	fmt.Println(info)
	fmt.Println()
}

/**
 * 显示群聊信息
 */
func outputGroupMes(mes *message.Message) { //mes 一定是smsMes类型
	//显示即可
	//1.反序列化mes.Data
	var smsResMes message.SmsResMes
	err := json.Unmarshal([]byte(mes.Data), &smsResMes)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err.Error())
		return
	}

	//显示信息
	info := fmt.Sprintf("【群聊】%s:\t%s", smsResMes.UserName, smsResMes.Content)
	fmt.Println(info)
	fmt.Println()
}
