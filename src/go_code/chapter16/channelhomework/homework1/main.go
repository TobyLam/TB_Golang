package main

import "fmt"

/*
* 要求：
1） 启动一个协程，将1-2000的数写入一个channel中，比如numChan
2）启动8个协程，从numChan取出数（比如n），并计算1+..+n的值，并存放到resChan
3) 最后8个协程协调完成工作后，再遍历resChan，显示结果【如res[1]=1...res[10]=55..】
$) 注意：考虑resChan chan int 是否合适
*/

func readData(numChan chan int, resChan chan map[int]int, exitChan chan bool) {

	for {
		n, ok := <-numChan
		if !ok {
			break
		}

		res := 0
		for i := 1; i <= n; i++ {
			res += i
		}

		resChan <- map[int]int{n: res}
	}

	exitChan <- true
}

func writeData(numChan chan int) {
	for i := 1; i <= 2000; i++ {
		numChan <- i
		//fmt.Println("writeData ", i)
	}
	close(numChan)
}

func main() {

	//1）将数写入协程
	numChan := make(chan int, 2000)
	go writeData(numChan)

	//启动八个协程,从numChan取出数，并计算
	resChan := make(chan map[int]int, 2000)

	//执行结果记录
	exitChan := make(chan bool, 8) //用来判断所有协程对于resChan写入操作是否完成

	for i := 1; i <= 8; i++ {
		go readData(numChan, resChan, exitChan)
	}

	for {
		if len(exitChan) > 7 {
			close(resChan)
			close(exitChan)
			break
		}
	}

	//遍历resChan
	for v := range resChan {
		for key, val := range v {
			fmt.Printf("res[%v]=%v\n", key, val)
		}

	}

}
