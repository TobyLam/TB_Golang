package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go 向redis写入数据和读取数据
	//1.链接到redis
	conn, err := redis.Dial("tcp", "192.168.1.18:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}

	defer conn.Close() //关闭...

	//密码认证
	_, err = conn.Do("Auth", "123456")
	if err != nil {
		fmt.Println("Auth err=", err)
		return
	}

	//2.通过go 向redis写入数据 sting [key-val]
	_, err = conn.Do("hmset", "user02", "name", "john", "age", 19)
	if err != nil {
		fmt.Println("hmset err=", err)
		return
	}

	//3.通过go 向redis读取数据 string [key-val]
	r, err := redis.Strings(conn.Do("hmget", "user02", "name", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	//fmt.Printf("r=%v type=%T\n", r, r)
	for i, v := range r {
		fmt.Printf("r[%d]=%s\n", i, v)
	}
}
