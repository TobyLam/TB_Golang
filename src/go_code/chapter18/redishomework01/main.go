package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

/*
*
要求：
1.Monster信息[name,age,skil]
2.通过终端输入三个monster的信息，使用golang操作redis,存放到redis中【比如使用hash数据类型】
3.编程，遍历出所有的monster信息,并显示在终端
4.提示：保存monster可以使用hash数据类型,遍历时先取出所有的keys 比如 keys monster*
*/
func main() {

	name := ""
	age := 0
	skill := ""

	fmt.Printf("name=")
	fmt.Scanln(&name)
	fmt.Printf("age=")
	fmt.Scanln(&age)
	fmt.Printf("skill=")
	fmt.Scanln(&skill)

	conn, err := redis.Dial("tcp", "192.168.1.18:6379")
	if err != nil {
		fmt.Println("Conn err=", err)
		return
	}

	defer conn.Close()

	_, err = conn.Do("Auth", "123456")
	if err != nil {
		fmt.Println("Redis auth err=", err)
		return
	}

	_, err = conn.Do("Hmset", "Monster", "name", name, "age", age, "skill", skill)
	if err != nil {
		fmt.Println("Hmset err=", err)
		return
	}

	Monster, err := redis.Strings(conn.Do("hmget", "Monster", "name", "age", "skill"))
	if err != nil {
		fmt.Println("Hmget err=", err)
	}

	//fmt.Printf("Monster=%v Type=%T", Monster, Monster)
	for i, v := range Monster {
		fmt.Printf("Monster[%d]=%v\n", i, v)
	}

}
