package main

import (
	"fmt"
	"math/rand"
	_ "strconv"
)

//要求：
//1）创建一个Person结构体[Name,Age,Address]
//2)使用rand方法配合随机创建10个Person实例，并放入到channel中，
//3）遍历channel,将各个Person实例的信息显示在终端

type Person struct {
	Name    string
	Age     int
	Address string
}

func main() {

	//创建一个Person类型的管道
	pChan := make(chan Person, 10)
	var p Person

	for i := 0; i < 10; i++ {
		p.Name = fmt.Sprintf("person|%d", rand.Intn(10))
		p.Age = rand.Intn(100)
		p.Address = "china"
		pChan <- p
	}
	close(pChan)

	for v := range pChan {
		fmt.Printf("姓名：%v 年龄：%v 地址：%v\n", v.Name, v.Age, v.Address)
	}
}
