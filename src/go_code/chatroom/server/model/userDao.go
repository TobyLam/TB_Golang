package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom/common/message"
)

// 在启动服务器启动后，就初始化一个userDao实例
// 全局变量，在需要redis操作时，直接使用
var (
	MyUserDao *UserDao
)

// 定义一个UserDao结构体
// 完成对User结构体的各种操作
type UserDao struct {
	pool *redis.Pool
}

// 使用工厂模式，创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

// 在UserDao 提供的方法
// 1.根据用户id 返回 一个 User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {

	//通过给定id 去 redis查询这个用户
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		//错误
		if err == redis.ErrNil { //表示在 users 哈希中，没有找到对应id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//需要把res反序列化为User实例
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

// 完成登录校验 Login
// 1. Login完成对用户的验证
// 2. 如果用户的id和pwd都正确，则返回一个user实例
// 3. 如果用户的id或pwd错误，则返回对应的错误信息
func (this *UserDao) Login(userId int, userPwd string) (user *User, err error) {

	//先从UserDao 的连接池中取一个连接
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	//根据userId查询用户信息
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//获取到用户信息后，校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

// 完成用户注册
func (this *UserDao) Register(user *message.User) (err error) {

	//先从UserDao 的连接池中取一个连接
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	//根据userId查询用户信息
	_, err = this.getUserById(conn, user.UserId)
	//用户id已存在，返回
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//用户id不存在
	data, err := json.Marshal(user) //序列化
	if err != nil {
		return
	}

	//入库
	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
		return
	}
	return
}

/**
 * 获取所有注册用户id
 */
func (this *UserDao) UserIds() (userIdsSlice []int, err error) {
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	userIdsSlice = make([]int, 0)

	res, err := redis.Ints(conn.Do("hkeys", "users"))
	if err != nil {
		fmt.Println("userDao UserIds() conn.Do(hkeys) err=", err)
		return
	}
	// 追加数据
	userIdsSlice = append(userIdsSlice, res...)

	return
}
