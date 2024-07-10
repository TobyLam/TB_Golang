package main

import (
	"fmt"
)

func main() {

	intChan := make(chan int, 100)

	intChan <- 100
	intChan <- 200

	close(intChan) //close
	//这是不能再写入数据到channel
	//intChan <- 300
	fmt.Println("okokk~")
	//当管道关闭后，仍可以读取数据
	n1 := <-intChan
	fmt.Println("n1=", n1)

	//遍历管道
	intChan2 := make(chan int, 100)
	for i := 0; i < 100; i++ {
		intChan2 <- i * 2 //放入100个数据到管道
	}

	//遍历,不能使用普通的for循环
	//长度会发生变化
	//for i := 0; i < len(intChan2); i++ {
	//
	//}

	//在遍历时，如果channel没有关闭，则会出现deadlock的错误
	//在遍历时，如果channel已经关闭，则会正常遍历数据，遍历完后，退出遍历
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v=", v)
	}

}
