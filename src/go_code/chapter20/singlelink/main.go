package main

import (
	"fmt"
)

// 定义一个HeroNode
type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode //表示指向下一个节点
}

// 给链表插一个节点
// 第一种插入方式，在单链表的最后加入
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.先找到链表的最后节点
	//2.创建一个辅助节点
	temp := head
	for {
		if temp.next == nil { //表示找到最后节点
			break
		}
		temp = temp.next //让temp不断指向下一个节点
	}

	//3.将newHeroNode加入链表最后
	temp.next = newHeroNode
}

// 给链表插一个节点
// 第二种插入方式，根据no 的编号从小到大插入..
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路
	//1.先找到链表的适当节点
	//2.创建一个辅助节点
	temp := head
	flag := true
	//比较插入节点的no 和 temp的下一个节点的no
	for {
		if temp.next == nil { //说明到链表最后
			break
		} else if temp.next.no >= newHeroNode.no { //若允许no相等，改为 >=
			//说明newHeroNode 就应该插入到temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明链表中已存在这个no,不允许插入
			flag = false
			break
		}

		temp = temp.next
	}

	if !flag {
		fmt.Println("已存在no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next
		temp.next = newHeroNode
	}

}

/**
 * 删除一个节点
 */
func DelHeroNode(head *HeroNode, id int) {
	temp := head
	flag := false
	//找到删除节点的no，和temp的下一个节点的no比较
	for {
		if temp.next == nil { //说明到链表的最后
			break
		} else if temp.next.no == id {
			// 说明找到了该节点
			flag = true
			break
		}
		temp = temp.next
	}

	if flag { //找到，删除
		temp.next = temp.next.next
	} else {
		fmt.Println("节点不存在...")
	}
}

// 显示链表的所有节点信息
func ListHeroNode(head *HeroNode) {
	//1.创建一个辅助节点
	temp := head

	//先判断该链表是不是一个空的链表
	if temp.next == nil {
		fmt.Println("空链表...")
		return
	}

	//2. 遍历链表
	for {
		fmt.Printf("[%d , %s , %s]==>", temp.next.no, temp.next.name, temp.next.nickname)
		//判断是否链表最后节点
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
}

func main() {

	//1.先创建一个头节点
	head := &HeroNode{}

	//2.创建一个新的HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "吴用",
		nickname: "智多星",
	}

	/*hero4 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}*/

	//3.加入
	InsertHeroNode2(head, hero3)
	InsertHeroNode2(head, hero1)
	InsertHeroNode2(head, hero2)
	//InsertHeroNode2(head, hero4)

	//4.显示
	ListHeroNode(head)

	//5 删除
	fmt.Println()
	DelHeroNode(head, 1)
	DelHeroNode(head, 3)
	ListHeroNode(head)
}
