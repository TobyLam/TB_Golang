package main

import (
	"errors"
	"fmt"
	"os"
)

// 使用结构体管理队列
type Queue struct {
	maxSize int
	array   [5]int // 数组=>模拟队列
	front   int    //表示指向队首(不包含第一个元素)
	rear    int    //表示指向队尾(包含最后的元素)
}

/**
 * 添加数据到队列
 */
func (this *Queue) AddQueue(val int) (err error) {

	//先判断队列是否已满
	if this.rear == this.maxSize-1 {
		return errors.New("queue full")
	}

	this.rear++ //rear后移
	this.array[this.rear] = val

	return
}

/**
 * 从队列中取出数据
 */
func (this *Queue) GetQueue() (val int, err error) {

	//先判断队列是否已空
	if this.rear == this.front { //对空
		return -1, errors.New("queue empty")
	}

	this.front++ //front++
	val = this.array[this.front]

	return
}

/**
 * 显示队列,遍历队首->队尾
 */
func (this *Queue) ShowQueue() {

	fmt.Println("当前队列：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d]=%d\t", i, this.array[i])
	}
	fmt.Println()
}

func main() {
	//先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err.Error())
			} else {
				fmt.Println("加入队列成功")
			}
		case "show":
			queue.ShowQueue()
		case "get":
			val, err := queue.GetQueue()
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
