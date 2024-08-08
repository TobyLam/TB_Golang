package message

// 定义一个用户的结构体
type User struct {
	//字段
	//为了序列化和反序列化成功
	//用户信息的字符串的key 和 结构体的字段对应的tag 名字一样！！！
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
