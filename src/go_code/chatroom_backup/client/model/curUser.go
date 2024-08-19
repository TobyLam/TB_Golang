package model

import (
	"go_code/chatroom/common/message"
	"net"
)

// 因为在客户端，很多地方会使用到curUser，作为全局变量
type CurUser struct {
	Conn net.Conn
	message.User
}
