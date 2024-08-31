package process2

import (
	"encoding/json"
	"fmt"
	"go_code/chatroom/common/message"
	"go_code/chatroom/server/model"
	"go_code/chatroom/server/utils"
	"net"
)

type UserProcess struct {
	//字段？
	Conn net.Conn //连接
	//新增一个字段，表示该Conn是哪个用户
	UserId int
}

/**
 * 区分状态变化消息 的 状态 ，分别处理相应事务
 */
func (this *UserProcess) DealUserStatus(mes *message.Message) {
	// 声明一个用户状态变化消息实例
	notifyUserStatusMes := &message.NotifyUserStatusMes{}
	// 反序列化mes,并赋值给notifyUserStatusMes
	err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
	if err != nil {
		fmt.Println("DealUserStatus json.Unmarshal() err=", err)
		return
	}

	// 判断用户是否离线
	switch notifyUserStatusMes.Status {
	case message.UserOffline:
		// 删除在线用户
		userMgr.DelOnlineUser(notifyUserStatusMes.UserId)
		// 通知其他用户 有用户离线
		this.NotifyOthersOfflineUser(notifyUserStatusMes)
	default:
		fmt.Println("暂不处理 非离线 状态业务...")
	}
}

/**
 * 通知在线用户有用户离线
 */
func (this *UserProcess) NotifyOthersOfflineUser(notifyUserStatusMes *message.NotifyUserStatusMes) {
	//遍历 onlineUsers,逐个发送 NotifyUserStatusMes
	for _, up := range userMgr.onlineUsers {
		//过滤当前用户
		//if id == userId {
		//	continue
		//}
		//通知
		this.NotifyMeOffline(notifyUserStatusMes, up.Conn)
	}
}

/**
 * 通知单个在线用户 有用户离线
 */
func (this *UserProcess) NotifyMeOffline(notifyUserStatusMes *message.NotifyUserStatusMes, conn net.Conn) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	//序列化notifyUserStatusMes
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal() err=", err)
		return
	}
	//将序列化的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//序列化mes,准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal() err=", err)
		return
	}

	// 创建 Transfer实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOffline err=", err)
		return
	}

}

// 通知所有在线用户
// userId 通知其他的在线用户，上线
func (this *UserProcess) NotifyOthersOnlineUser(userId int) {

	//遍历 onlineUsers,逐个发送 NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		//过滤当前用户
		if id == userId {
			continue
		}
		//通知
		up.NotifyMeOnline(userId)
	}
}

/**
 * 通知单个在线用户有用户在线
 */
func (this *UserProcess) NotifyMeOnline(userId int) {
	//组装NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marsha1 err=", err)
		return
	}
	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对mes再次序列化，准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err=", err)
		return
	}

	//发送,创建一个Transfer实例，发送
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err=", err)
		return
	}
}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	//1.从mes中取出 mes.Data,并直接反序列化成registerMes
	var registerMes message.RegisterMes
	json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType

	//2再声明一个RegisterResMes，并完成赋值
	var registerResMes message.RegisterResMes

	// 到redis数据库去完成注册
	//1. 使用model.MyUserDao去redis验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生未知错误..."
		}
	} else {
		registerResMes.Code = 200
	}

	//3将 loginResMes 序列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6.发送data,封装到writePkg函数
	//因为使用了分层模式(mvc)，先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return

}

// 编写一个函数serverProcessLogin函数，专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	//核心代码...
	//1.先从mes中取出 mes.Data ,并直接反序列化成LoginMes
	var loginMes message.LoginMes
	json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err=", err)
		return
	}

	//1先声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2再声明一个LoginResMes，并完成赋值
	var loginResMes message.LoginResMes

	// 到redis数据库去完成验证
	//1. 使用model.MyUserDao去redis验证
	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)

	if err != nil {

		if err == model.ERROR_USER_NOTEXISTS {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}

	} else {
		loginResMes.Code = 200
		loginResMes.UserName = user.UserName
		//这里因为用户登录成功，把该登录成功的用户放入到userMgr中
		//将登录成功的用户的userId 赋给 this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)
		//通知其他的在线用户，上线了
		this.NotifyOthersOnlineUser(loginMes.UserId)
		//将当前在线用户的id 放入到loginResMes.UsersId
		//遍历 userMgr.onlineUsers
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UsersId = append(loginResMes.UsersId, id)
		}
		fmt.Println(user, "登录成功")
	}

	//3将 loginResMes 序列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//4.将data赋值给resMes
	resMes.Data = string(data)

	//5.对resMes进行序列化，准备发送
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6.发送data,封装到writePkg函数
	//因为使用了分层模式(mvc)，先创建一个Transfer实例，然后读取
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)

	//发送离线信息
	if loginResMes.Code == 200 {
		//遍历离线用户列表切片
		//判断是否存在未发送的离线信息
		for sendUserId, intSlice := range userMgr.offlineUserIds {
			for index, offlineId := range intSlice {
				if offlineId != loginMes.UserId {
					//非当前登录成功用户
					continue
				} else {
					//存在离线留言
					sms, _ := model.MySmsDao.GetSms(sendUserId)
					//sms中的ReceiveId为空，赋值当前登录成功的用户id
					sms.ReceiverId = loginMes.UserId

					//创建一个smsProcess实例
					smsProcess := &SmsProcess{}
					err = smsProcess.SendMesToEachOnlineUser(sms, userMgr.onlineUsers[loginMes.UserId].Conn)
					if err == nil {
						//发送成功
						// 将offlineUserIds中的该id删除
						// 由于在切片中删除较困难，选择将其值设置为-1
						intSlice[index] = -1
					} else {
						fmt.Println("用户", loginMes.UserId, "留言信息发送失败")
					}

					// MyUserMgr.offlineUserIds 的 intslice 标识位自动减一
					intSlice[len(intSlice)-1] = intSlice[len(intSlice)-1] - 1
					// 跳出该轮遍历
					break
				}
			}

			// 判断 MyUserMgr.offlineUserIds 的 intslice 标识位是否为0
			if intSlice[len(intSlice)-1] == 0 {
				// 删除数据库中缓存的信息
				model.MySmsDao.DelSms(sendUserId)
				// fmt.Println("缓存删除")
			}
		}
	}

	return
}
