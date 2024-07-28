package main

import (
	"fmt"
	"gopkg.in/redis.v4"
)

func createClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.18:6379",
		Password: "123456",
		DB:       0,
	})

	//通过client.Ping（）来检查是否成功连接到redis服务器
	pong, err := client.Ping().Result()

	fmt.Println(pong, err)

	return client
}

func main() {

	client := createClient()

	defer client.Close()

	val, err := client.Get("name").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("name=", val)

}
