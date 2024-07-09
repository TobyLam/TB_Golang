package main

import (
	"fmt"
	"sync"
	_ "time"
)

//需求，现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
//最后显示出来。要求使用goroutine完成

// 思路
// 1.编写一个函数，来计算各个数的阶乘，并放入到map中
// 2.启动多个协程，统一将结果放入的map中
// 3.map应该做成全局的
var (
	myMap = make(map[int]int, 10)
	//声明一个全局的互斥锁
	//lock 是一个全局的互斥锁，
	//sync 是包：synchronized 同步
	//Mutex 是互斥
	lock sync.Mutex
)

// test函数计算 n!,将这个结果放入到myMap
func test(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	//将res放入到myMap
	//加锁
	lock.Lock()
	myMap[n] = res //concurrent map writes?
	//解锁
	lock.Unlock()
}

func main() {
	//开启协程完成这个任务
	for i := 1; i <= 20; i++ {
		go test(i)
	}

	//休眠10秒钟 ？ [等待时间如何设置]
	//time.Sleep(time.Second * 5)

	//输出结果(遍历)
	lock.Lock()
	for i, v := range myMap {
		fmt.Printf("map[%d]=%d\n", i, v)
	}
	lock.Unlock()
}
