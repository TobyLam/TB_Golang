package main

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// 定义一个全局的pool
var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {

	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) { //初始化连接，redis服务器ip地址
			c, err := redis.Dial("tcp", address)
			if err != nil {
				return nil, err
			}
			if _, err = c.Do("Auth", "123456"); err != nil {
				c.Close()
				return nil, err
			}

			return c, err
		},
		MaxIdle:     maxIdle,     //最大空闲连接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数， 0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
	}
}
