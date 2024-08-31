package process2

import (
	"fmt"
)

// 因为UserMgr 实例在服务器有且仅有一个
// 且在多处使用，因此定义为全局变量
var (
	userMgr *UserMgr
)

type UserMgr struct {
	onlineUsers    map[int]*UserProcess //在线用户，[当前登录用户id]关联的用户处理实例（包含链接和当前用户id，方便后续操作）
	offlineUserIds map[int][]int        //离线用户id的map
}

// userMgr初始化
func init() {
	userMgr = &UserMgr{
		onlineUsers:    make(map[int]*UserProcess, 1024), //初始化，长度1024
		offlineUserIds: make(map[int][]int, 1024),        //初始化，长度为0
	}
}

// 完成对onlineUsers添加
func (this *UserMgr) AddOnlineUser(up *UserProcess) {
	this.onlineUsers[up.UserId] = up
}

// 删除
func (this *UserMgr) DelOnlineUser(userId int) {
	delete(this.onlineUsers, userId)
}

// 返回当前所有在线用户
func (this *UserMgr) GetAllOnlineUser() map[int]*UserProcess {
	return this.onlineUsers
}

// 根据id返回对应的值
func (this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess, err error) {
	//如何从map取出一个值，带检测方式
	up, ok := this.onlineUsers[userId]
	if !ok { //说明，需要查找的这个用户不在线
		err = fmt.Errorf("用户%d 不存在", userId)
		return
	}
	return
}
