package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用一个结构体管理环形队列
type CircleQueue struct {
	maxSize int    //5
	array   [5]int //数组
	head    int    //指向队首，初值0,指向第一个元素
	tail    int    //指向队尾，初值0，不包含最后一个元素
}

// 入/出队列方法 AddQueue(push) GetQueue(pop)
// 入队列
func (this *CircleQueue) Push(val int) (err error) {
	if this.IsFull() {
		return errors.New("queue full")
	}
	this.array[this.tail] = val //把值给尾部
	this.tail = (this.tail + 1) % this.maxSize

	return
}

// 出队列
func (this *CircleQueue) Pop() (val int, err error) {
	if this.IsEmpty() {
		return 0, errors.New("queue empty")
	}
	//取出
	val = this.array[this.head]
	this.head = (this.head + 1) % this.maxSize
	return
}

// 显示队列
func (this *CircleQueue) ListQueue() {
	fmt.Println("环形队列情况如下：")
	//取出当前队列有多少个元素
	size := this.Size()
	if size == 0 {
		fmt.Println("队列为空")
	}

	//辅助变量，指向head
	tmpHead := this.head
	for i := 0; i < size; i++ {
		fmt.Printf("arr[%d]=%d\n", tmpHead, this.array[tmpHead])
		tmpHead = (tmpHead + 1) % this.maxSize
	}
	fmt.Println()
}

// 队满
func (this *CircleQueue) IsFull() bool {
	return (this.tail+1)%this.maxSize == this.head
}

// 对空
func (this *CircleQueue) IsEmpty() bool {
	return this.tail == this.head
}

// 队长(元素个数)
func (this *CircleQueue) Size() int {
	//关键算法
	return (this.tail - this.head + this.maxSize) % this.maxSize
}

func main() {

	//初始化一个环形队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	var key string
	var val int

	for {
		fmt.Println("1.输入add 表示添加数据到队列")
		fmt.Println("2.输入get 表示从队列获取数据")

		fmt.Println("3.输入show 表示显示队列")
		fmt.Println("4.输入exit 表示退出")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("请输入入队数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "show":
			queue.ListQueue()
		case "get":
			val, err := queue.Pop()
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("从队列中取出了一个数=", val)
			}
		case "exit":
			os.Exit(0)
		}
	}
}
