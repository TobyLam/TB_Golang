package message

// 定义消息类型常量
// 登录消息、登录响应消息
const (
	LoginMesType            = "LoginMes"            //登录
	LoginResMesType         = "LoginResMes"         //登录响应
	RegisterMesType         = "RegisterMes"         //注册
	RegisterResMesType      = "RegisterResMes"      //注册响应
	NotifyUserStatusMesType = "NotifyUserStatusMes" //用户状态改变
	SmsMesType              = "SmsMes"              //发送消息
)

// 定义用户状态的常量
const (
	UserOnline = iota
	UserOffline
	UserBusyStatus
)

// 定义消息结构体，包含消息类型和数据
type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息的内容
}

// 根据需求定义各类消息结构体，可扩展

// 登录请求消息
type LoginMes struct {
	UserId   int    `json:"userId"`   //用户id
	UserPwd  string `json:"userPwd"`  //用户密码
	UserName string `json:"userName"` //用户名
}

// 登录响应消息
type LoginResMes struct {
	Code    int    `json:"code"` //返回状态码 500 表示该用户未注册 200表示登录成功
	UsersId []int  //增加字段，保存用户id的切片
	Error   string `json:"error"` //返回错误信息
}

// 注册请求消息
type RegisterMes struct {
	User User `json:"user"` //类型就是User结构体
}

// 注册响应消息
type RegisterResMes struct {
	Code  int    `json:"code"`  //返回状态码 400 表示该用户已存在 200表示注册成功
	Error string `json:"error"` //返回错误信息
}

// 配合服务器端推送用户状态变化的消息
type NotifyUserStatusMes struct {
	UserId int `json:"userId"` //用户id
	Status int `json:"status"` //用户的状态
}

// 增加一个SmsMes //发送的消息
type SmsMes struct {
	Content string `json:"content"` //内容
	User           //匿名结构体
}
