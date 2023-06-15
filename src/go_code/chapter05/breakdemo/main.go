package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//为了生成一个随机数，还需要给rand设置一个种子
	//time.Now().Unix() :返回一个秒数，时间戳
	//rand.Seed(time.Now().Unix())
	//如何随机的生成1-100整数
	//n := rand.Intn(100) + 1 //[0 100)
	//fmt.Println(n)

	//随机生成 1-100 的一个数，直到生成了 99 这个数，看看你一共用了几次？
	// 分析： 编写一个无限循环的控制，然后不停的生成随机数，当生成了99时，就退出这个无限循环==>break
	var count int = 0
	for {
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100) + 1
		count++
		if n == 99 {
			break
		}
	}
	fmt.Printf("生成99一共使用了%d次 \n", count)

	//指定标签的形式使用 break
lable2:
	for i := 0; i < 4; i++ {
		//lable1: //设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break lable1 // break 默认跳出最近的for循环
				break lable2
			}
			fmt.Println("j=", j)
		}
	}
}
