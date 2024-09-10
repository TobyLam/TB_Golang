package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
*
  - 要求：
  - 1）创建一个数组模拟队列，每隔一定时间（随机的）,给该数组添加一个数。
  - 2）启动两个协程，每隔一定时间（时间随机）到队列取出数据
  - 3) 在控制台输出
    x号协程 服务---》x号客户
    x号协程 服务---》x号客户
    ...
  - 4) 使用锁机制即可
*/

// 创建一个结构体，实现数组模拟队列
type ArrayQueue struct {
	array   [50]int //数据
	maxSize int     //队长
	head    int     //指向队首
	tail    int     //指向队尾
}

// 声明一个全局互斥锁
var lock sync.Mutex

// 入队列
func (this *ArrayQueue) Push(val int) {
	//判断是否队满
	if this.IsFull() {
		fmt.Println("队列满了...")
		return
	}

	//入队
	this.array[this.tail] = val
	this.tail = (this.tail + 1) % this.maxSize //队尾加一，防止超出最大长度，进行取余
	fmt.Printf("%d号客户排队了\n", val)
}

/**
 * 出队列
 * param n int 表示协程编号
 */
func (this *ArrayQueue) Pop(n int) {
	for {
		//判断是否队空
		if this.IsEmpty() {
			//休眠1秒
			time.Sleep(time.Second)
			fmt.Println("没有客户排队...")
		} else {
			//随机隔一段时间取出数据
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

			//加锁，避免协程资源竞争
			lock.Lock()

			//取出一个数据
			val := this.array[this.head]
			fmt.Printf("%d号协程 服务---》%d号客户\n", n, val)
			//队列头部后移
			this.head = (this.head + 1) % this.maxSize

			//解锁
			lock.Unlock()
		}
	}
}

// 判断队满
func (this *ArrayQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 判断队空
func (this *ArrayQueue) IsEmpty() bool {
	return this.tail == this.head
}

func main() {
	//初始化一个环形队列
	queue := &ArrayQueue{
		maxSize: 50,
		head:    0,
		tail:    0,
	}

	//随机种子
	rand.Seed(time.Now().UnixNano())

	//定义一个随机数，准备入队
	var num int

	//启动两个协程，读取数据
	go queue.Pop(1)
	go queue.Pop(2)

	//每个一段时间，将num进行入队
	for {
		num = rand.Intn(100)
		//fmt.Printf("客户%d准备排队\n", num)
		//随机休眠一段时间
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		//入队
		queue.Push(num)
	}
}
