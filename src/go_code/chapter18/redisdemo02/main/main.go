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
	_, err = conn.Do("hset", "user01", "name", "john")
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	_, err = conn.Do("hset", "user01", "age", "10")
	if err != nil {
		fmt.Println("hset err=", err)
		return
	}

	//3.通过go 向redis读取数据 string [key-val]
	r1, err := redis.String(conn.Do("hget", "user01", "name"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	r2, err := redis.Int(conn.Do("hget", "user01", "age"))
	if err != nil {
		fmt.Println("hget err=", err)
		return
	}

	//因为返回 r是 interface{}
	//因为name 对应的值是string,因此需要转换
	//nameString := r.(string) //error

	fmt.Printf("操作ok r1=%v r2=%v\n", r1, r2)
}
