package main

import (
	"fmt"
)

//约瑟夫问题

// 节点结构体
type Node struct {
	No   int   //编号
	Next *Node //指向下一个节点,默认值 nil
}

/*
 * 编写一个函数，构成单向环形链表
 * param num:节点个数
 * return *Node: 返回该环形链表的头节点
 */
func AddNode(num int) *Node {

	first := &Node{}   //空节点
	curNode := &Node{} //空节点

	//判断
	if num < 1 {
		fmt.Println("num的值不正确")
		return first
	}

	//循环构建环形链表
	for i := 1; i <= num; i++ {
		node := &Node{
			No:   i,
			Next: nil,
		}

		if i == 1 { //头结点【特殊处理】
			first = node //头结点不能变动
			curNode = node
			curNode.Next = first //形成环形链表
		} else {
			curNode.Next = node  //当前节点指向新节点
			curNode = node       //当前节点往后移，新节点成为当前节点
			curNode.Next = first //新节点指向头节点，构成环形链表
		}
	}

	return first
}

// 遍历
func ListNode(first *Node) {

	//处理空链表
	if first.Next == nil {
		fmt.Println("空链表...")
		return
	}

	//创建一个辅助节点
	curNode := first

	for {
		fmt.Printf("节点编号=%d ->", curNode.No)
		//退出条件
		if curNode.Next == first {
			break
		}
		//后移
		curNode = curNode.Next
	}
}

// 链表长度
func GetLen(first *Node) (len int) {
	if first.Next == nil {
		return
	}
	curNode := first
	for {
		len++
		if curNode.Next == first {
			break
		}
		curNode = curNode.Next
	}

	return
}

/**
 * Josephu 问题为：设编号为 1，2，… n 的 n 个人围坐一圈，约定编号为 k（1<=k<=n）的人从 1
 * 开始报数，数到 m 的那个人出列，它的下一位又从 1 开始报数，数到 m 的那个人又出列，依次类推，
 * 直到所有人出列为止，由此产生一个出队编号的序列。
 */
//思路
//1.编写一个函数，Josephu(first *Node,startNo int,countNum int)
//2.使用一个算法，按照要求，在环形链表中留下最后一个人
func JosePhu(first *Node, startNo, countNum int) {

	//处理空链表
	if first.Next == nil {
		fmt.Println("空链表...")
		return
	}

	//判断StartNo
	if startNo > GetLen(first) {
		fmt.Println("初始编号大于节点个数...")
		return
	}

	// 定义一个辅助节点，帮助删除节点
	tail := first
	// 让tail指向链表最后一个节点,使first和tail保持前后节点
	for {
		if tail.Next == first {
			break
		}
		tail = tail.Next
	}

	// 让first 移动到 startNo
	for i := 1; i <= startNo-1; i++ {
		first = first.Next
		tail = tail.Next
	}

	// 开始数countNum次，然后删除first 指向的节点
	for {
		//数countNum次,first和tail节点同时移动
		for i := 1; i <= countNum-1; i++ {
			first = first.Next
			tail = tail.Next
		}
		fmt.Printf("编号%d 的节点出队-》\n", first.No)
		//删除first指向的节点
		first = first.Next
		tail.Next = first

		//判断只剩一个节点,退出循环
		if tail == first {
			break
		}
	}
	fmt.Printf("最后出队的节点编号%d->\n", first.No)

}

func main() {

	first := AddNode(500)

	//显示
	ListNode(first)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	JosePhu(first, 20, 31)

}
