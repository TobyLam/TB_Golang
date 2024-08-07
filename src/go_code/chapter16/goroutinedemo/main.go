package main

import (
	"fmt"
	"strconv"
	"time"
)

//1) 在主线程(可以理解成进程)中，开启一个 goroutine, 该协程每隔 1 秒输出 "hello,world"
//2) 在主线程中也每隔一秒输出"hello,golang", 输出 10 次后，退出程序
//3) 要求主线程和 goroutine 同时执行.

// 编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("test() hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启了一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println("main() hello,world" + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}
