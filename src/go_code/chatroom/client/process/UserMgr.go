package process

import (
	"fmt"
	"go_code/chatroom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)

// 在客户端显示当前在线的用户
func outputOnlineUser() {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id：\t", id)
	}
}

// 编写方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//优化
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	if !ok { //原来没有
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	//已存在，则直接更新状态
	user.UserStatus = notifyUserStatusMes.UserId

	onlineUsers[notifyUserStatusMes.UserId] = user

	outputOnlineUser()
}
