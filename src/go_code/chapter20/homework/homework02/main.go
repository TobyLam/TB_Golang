package main

import (
	"errors"
	"fmt"
)

/**
 * 要求：
 * 1） 创建一个链表模拟队列，实现数据入队列，出队列，显示队列
 */

//思路：从head端插入数据，从tail端删除数据

// 链表实现队列的方法定义

// 定义一个节点结构体
type node struct {
	Item interface{} //可赋值任何类型的数据
	Next *node       //下一个节点
}

// 定义一个链表队列结构体
type linkedListQueue struct {
	Length int   //队列长度（实际节点数）
	head   *node //头节点
	tail   *node //尾节点
}

// 队列是否为空
func (this *linkedListQueue) IsEmpty() bool {
	//return this.Length == 0
	return this.head == nil
}

// 队列长度
func (this *linkedListQueue) Len() int {
	return this.Length
}

// 入队
func (this *linkedListQueue) Enqueue(item interface{}) {
	// 创建一个节点
	temp := &node{
		Item: item,
		Next: nil,
	}

	if this.Length == 0 { //如果队列为空，则直接入队，队首、队尾均指向此节点
		this.tail = temp
		this.head = temp
	} else { //队列不为空，则在队尾插入此节点
		this.tail.Next = temp      //在队尾插入节点
		this.tail = this.tail.Next //队尾后移
	}
	this.Length++ //队列长度加1
}

// 出队
func (this *linkedListQueue) Dequeue() (item interface{}) {
	//判断队列是否为空（节点个数为0）
	if this.Length == 0 {
		return errors.New(
			"failed to dequeue, queue is empty")
	}

	item = this.head.Item      //从队首取出数据（先进先出）
	this.head = this.head.Next //头结点后移

	// 当只有一个元素时，出列后head和tail都等于nil
	// 这时要将tail置为nil，不然tail还会指向第一个元素的位置
	// 比如唯一的元素原本为2，不做这步tail还会指向2
	if this.Length == 1 {
		this.tail = nil
	}

	this.Length-- //队列长度（节点个数减一）
	return
}

/**
 * 将队列遍历，并格式化返回
 */
func (this *linkedListQueue) Traverse() (resp []interface{}) {

	if this.IsEmpty() {
		fmt.Println("空队列")
		return
	}

	temp := this.head
	for {
		if temp.Next == nil { //遍历到最后一个了
			resp = append(resp, temp.Item, "<--")
			break
		} else {
			resp = append(resp, temp.Item, "<--")
			temp = temp.Next //后移
		}
	}

	////遍历链表
	//for i := 0; i < this.Length; i++ {
	//	resp = append(resp, temp.Item, "<--")
	//	temp = temp.Next //后移
	//}
	return
}

/**
 * 获取头结点的数据
 */
func (this *linkedListQueue) GetHead() (item interface{}) {
	if this.Length == 0 {
		return errors.New(
			"failed to getHead, queue is empty")
	}
	return this.head.Item
}

// Back 获取该队列尾元素但不出队
func (this *linkedListQueue) GetTail() (item interface{}) {
	if this.IsEmpty() {
		fmt.Println("空队列")
		return nil
	}
	return this.tail.Item
}

func main() {

	//初始化一个链表队列
	queue := &linkedListQueue{}

	//入队五个数据
	for i := 0; i < 5; i++ {
		//入队，格式化显示
		queue.Enqueue(i)
		fmt.Println(queue.Traverse())
	}

	//当前队列的情况
	fmt.Printf("队首元素%d,队列长度%d,队列是否为空%v\n", queue.GetHead(), queue.Len(), queue.IsEmpty())

	for i := 0; i < 5; i++ {
		//出队，格式化显示
		queue.Dequeue()
		fmt.Println(queue.Traverse())
	}

	fmt.Printf("队首元素%d,队列长度%d,队列是否为空%v\n", queue.GetHead(), queue.Len(), queue.IsEmpty())
}
