package process

import (
	"fmt"
	"go_code/chatroom/client/model"
	"go_code/chatroom/common/message"
)

// 客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User, 10)
var CurUser model.CurUser //在用户登录成功后，完成对CurUser的初始化

// 在客户端显示当前在线的用户
func outputOnlineUser() (userNum int, err error) {
	//遍历onlineUsers
	fmt.Println("当前在线用户列表：")
	if userNum = len(onlineUsers); len(onlineUsers) == 0 {
		fmt.Println("其他用户都下线了...")
	} else {
		for id, _ := range onlineUsers {
			fmt.Println("用户id：\t", id)
		}
	}
	fmt.Println()
	return
}

// 编写方法，处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {

	//优化，不要直接赋值给onlineUsers，先判断是否存在，如果存在，直接修改状态即可
	user, ok := onlineUsers[notifyUserStatusMes.UserId]
	//如果原来没有此用户才新增
	if !ok {
		user = &message.User{
			UserId: notifyUserStatusMes.UserId,
		}
	}
	//已存在，则直接更新状态
	user.UserStatus = notifyUserStatusMes.Status

	//重新赋值
	onlineUsers[notifyUserStatusMes.UserId] = user

	//显示初当前用户外的在线用户列表
	//outputOnlineUser()
}
