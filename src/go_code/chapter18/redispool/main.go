package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

// 定义一个全局的pool
var pool *redis.Pool

// 当启动程序时，就初始化连接池
func init() {
	pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库最大连接数， 0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) { //初始化连接
			return redis.Dial("tcp", "192.168.1.18:6379")
		},
	}
}

func main() {
	//先从pool取出一个连接
	conn := pool.Get()
	defer conn.Close()

	conn.Do("Auth", "123456")

	_, err := conn.Do("set", "name", "汤姆猫~~~")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	//取出
	r, err := redis.String(conn.Do("get", "name"))
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}

	fmt.Println("r=", r)

	//如果要从pool取出数据，连接池不能关闭
	pool.Close()
	conn2 := pool.Get()

	_, err = conn2.Do("Set", "name2", "汤姆猫2")
	if err != nil {
		fmt.Println("conn.Do err=", err)
		return
	}
	//fmt.Println("conn2=", conn2)
}
