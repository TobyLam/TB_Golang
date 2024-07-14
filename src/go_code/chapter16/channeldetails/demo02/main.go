package main

import (
	"fmt"
	"time"
)

func main() {
	//使用select可以解决管道取数据的阻塞问题

	//1.定义一个管道，10个数据int
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	//2.定义一个管道，5个数据string
	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d", i)
	}

	//传统的方法在遍历管道时，如果不关闭会阻塞导致dead lock

	//问题，无法确定什么时候关闭管道
	//可以使用select 方法解决
	//label:
	for {
		//注意：这里，如果intChan一直没有关闭，不会一直阻塞导致deadlock
		//会字段到下一个case匹配
		select {
		case v := <-intChan:
			fmt.Printf("从intChan读取的数据%d\n", v)
			time.Sleep(time.Second)
		case v := <-stringChan:
			fmt.Printf("从stringChan读取的数据%s\n", v)
			time.Sleep(time.Second)
		default:
			fmt.Printf("都取不到，不玩了，程序员可以加入逻辑\n")
			time.Sleep(time.Second)
			//break label (结合label)
			return
		}
	}

}
