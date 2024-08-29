package process

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/client/utils"
	"go_code/chatroom/common/message"
)

type SmsProcess struct {
}

/*
 * 发送私聊的信息
 * @param content string 消息内容
 * @param receiverId int 接收对象id
 */
func (this *SmsProcess) SendPrivateMes(content string, receiverId int) (err error) {

	// 1.创建一个Mes实例
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.UserStatus
	smsMes.UserName = CurUser.UserName
	smsMes.ReceiverId = receiverId //指定接收者id

	// 3.序列化SmsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("json.Marshal(smsMes) err=", err)
		return
	}

	mes.Data = string(data)

	// 4.mes序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal(mes) err=", err)
		return
	}

	// 5. 将mes发送给服务器..
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	// 6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendPrivateMes err=", err.Error())
		return
	}
	return
}

/*
 * desc:发送群聊的消息
 * @param content string  消息内容
 */
func (this *SmsProcess) SendGroupMes(content string) (err error) {

	// 1.创建一个Mes实例
	var mes message.Message
	mes.Type = message.SmsMesType

	// 2.创建一个SmsMes 实例
	var smsMes message.SmsMes
	smsMes.Content = content               //内容.
	smsMes.UserId = CurUser.UserId         //发送者id
	smsMes.UserStatus = CurUser.UserStatus //发送者状态
	smsMes.UserName = CurUser.UserName     //发送者昵称

	// 3.序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	mes.Data = string(data)

	//4. 对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes json.Marshal fail =", err.Error())
		return
	}

	//5. 将mes发送给服务器..
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}

	//6.发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMes err=", err.Error())
		return
	}
	return

}
