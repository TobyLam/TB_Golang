package model

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"go_code/chatroom/common/message"
)

/**
 * Redis数据格式hash: UserId->Content
 */

/**
 * 创建SmsDao结构体
 * 用于对Sms增删改查
 */
type SmsDao struct {
	pool *redis.Pool
}

// 全局变量，在需要使用redis直接使用，
var (
	MySmsDao *SmsDao
)

// 使用工厂模式，创建一个SmsDao实例
func NewSmsDao(pool *redis.Pool) (smsDao *SmsDao) {
	smsDao = &SmsDao{
		pool: pool,
	}
	return
}

//SmsDao提供方法

/**
 * 新增
 */
func (this *SmsDao) SaveSms(smsMes *message.SmsMes) (err error) {
	//获取一个redis连接
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	//序列化smsMes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SaveSms json.Marshal() err=", err)
		return
	}

	//入库
	_, err = conn.Do("hset", "smses", smsMes.UserId, string(data))
	if err != nil {
		fmt.Println("SaveSms hset(smses) err=", err)
		return
	}

	return
}

/**
 * 用户离线消息全部接收后，删除
 */
func (this *SmsDao) DelSms(userId int) (err error) {
	//获取一个redis连接
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	_, err = conn.Do("hdel", "smses", userId)
	if err != nil {
		fmt.Println("DelSms hel(smses) err=", err)
		return
	}
	return
}

/**
 * 根据指定用户id获取离线消息
 */
func (this *SmsDao) GetSms(userId int) (sms *message.SmsMes, err error) {
	//获取一个redis连接
	conn := this.pool.Get()
	//延时关闭
	defer conn.Close()

	//获取
	res, err := redis.String(conn.Do("hget", "smses", userId))
	if err != nil {
		fmt.Println("GetSms() err=", err)
		return
	}
	err = json.Unmarshal([]byte(res), &sms)
	if err != nil {
		fmt.Println("GetSms json.Unmashal() err=", err)
		return
	}
	return
}
