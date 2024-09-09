package main

import (
	"fmt"
)

// 定义猫的结构体
type CatNode struct {
	no   int //编号
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {

	//判断是不是第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head //形成一个环状
		fmt.Println(newCatNode, "加入到环形的链表")
		return
	}

	//定义一个临时变量，找到环形链表的最后节点
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}

	//加入到环形链表
	temp.next = newCatNode
	newCatNode.next = head
}

// 输出环形链表
func ListCircleLink(head *CatNode) {
	fmt.Println("环形链表如下：")
	temp := head
	if temp.next == nil {
		fmt.Println("空环形链表...")
		return
	}
	for {
		fmt.Printf("猫的信息为[id=%d name=%s] ->\n", temp.no, temp.name)
		if temp.next == head {
			break
		}
		temp = temp.next
	}
}

// 删除一只猫
func DelCatNode(head *CatNode, id int) *CatNode {
	temp := head
	helper := head

	//空链表
	if temp.next == nil {
		fmt.Println("空环形链表...")
		return head
	}

	//如果只有一个节点
	if temp.next == head { //只有一个节点
		temp.next = nil
		fmt.Printf("头结点id=%d已删除\n", head.no)
		return head
	}

	//helper节点移动到链表最后
	for {
		if helper.next == head {
			break
		}
		helper = helper.next
	}

	//如果节点数量大于等于2
	flag := true
	for {
		if temp.next == head { //说明移动到最后一个节点,还没进行id比较（还没找到【未来得及比较】，退出再比较，否则难以退出for循环，退出for循环后再进行比较）
			break
		}
		if temp.no == id { //在最后一个之前找到了，立马进行删除，删除完成后退出for循环
			if temp == head { //删除头节点
				head = head.next //删除头节点后，必须返回新的头节点，main栈中的head需要更新
			}
			//找到
			helper.next = temp.next
			fmt.Printf("节点id=%d已删除\n", id)
			flag = false //id不重复才可以这样处理
			break
		}

		//temp和helper一直进行移动
		//helper先前处理过，移动到链表最后，则每循环一次，helper和temp始终保持指向前后节点关系
		temp = temp.next     //【比较是否是需要删除的节点】
		helper = helper.next // 【使用helper来删除temp】
	}

	if flag { //在for循环没有删除，遍历到最后一个节点还没来得及比较id和no，在此处进行比较
		if temp.no == id {
			helper.next = temp.next
			fmt.Printf("节点id=%d已删除\n", id)
		} else {
			fmt.Printf("没有这个节点,id=%d,无法删除\n", id)
		}
	}

	return head
}

func main() {

	//	初始化一个环形链表的头节点
	head := &CatNode{}

	//创建一只猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}
	cat2 := &CatNode{
		no:   2,
		name: "tom2",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom3",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)

	ListCircleLink(head)

	//删除一只猫,head节点可能被删除掉，重新指向另外的节点，需要重新接收
	head = DelCatNode(head, 30)
	head = DelCatNode(head, 1)
	//head = DelCatNode(head, 2)
	//head = DelCatNode(head, 3)
	//head = DelCatNode(head, 1)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	ListCircleLink(head)
}
