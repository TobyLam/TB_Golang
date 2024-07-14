package main

import (
	"fmt"
	"time"
)

/*
需求：
要求统计 1-200000 的数字中，哪些是素数？这个问题在本章开篇就提出了，现在我们有 goroutine
和 channel 的知识后，就可以完成了 [测试数据: 80000]
*/
func putNum(intChan chan int) {
	for i := 0; i < 200000; i++ {
		intChan <- i
	}

	//关闭
	close(intChan)
}

// 从 intChan取出数据，并判断是否为素数，如果是,就放入到primeChan
func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {

	//使用for循环
	var flag bool //标识是不是素数
	for {
		//time.Sleep(time.Millisecond * 10)
		num, ok := <-intChan
		if !ok { //intChan 取不到...
			break
		}
		flag = true //假设是素数
		//判断num是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { //说明该num不是素数
				flag = false
				break
			}
		}

		if flag {
			//将这个数放入到primeChan
			primeChan <- num
		}
	}

	fmt.Println("有一个primeNum 协程因为取不到数据，退出")
	//这里不能关闭primeChan

	//向exitChan 写入true
	exitChan <- true
}

func main() {

	intChan := make(chan int, 1000)
	primeChan := make(chan int, 200000) //存放结果
	exitChan := make(chan bool, 4)      //标识退出的管道

	start := time.Now().Unix()
	//开启一个协程，向 intChan放入 1-8000个数
	go putNum(intChan)

	//开启4个协程，从 intChan取出数据，并判断是否为素数，如果是,就放入到primeChan
	for i := 0; i < 4; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	//这里主线程进行处理
	//直接使用协程和匿名函数
	go func() {
		for i := 0; i < 4; i++ {
			<-exitChan //取不到4个true，会阻塞
		}
		end := time.Now().Unix()
		fmt.Println("使用协程耗时=", end-start)
		//当从exitChan取出了4个结果，就可以放心的关闭primeChan
		close(primeChan)
	}()

	//遍历primeChan，把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			break
		}
		//将结果输出
		//fmt.Printf("素数=%d\n", res)
	}

	fmt.Println("main线程退出")

}
