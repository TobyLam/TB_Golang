package main

import (
	"fmt"
	"os"
)

// 定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 方法待定
func (this *Emp) Info() {
	fmt.Printf("链表%d 找到该雇员 id=%d\n", this.Id%7, this.Id)
}

// 定义EmpLink
// 这里定义的EmpLink 不带表头，即第一个节点就存放雇员Emp
type EmpLink struct {
	Head *Emp
}

// 方法待定
// 1.添加员工的方法,保证添加时，编号从小到大
func (this *EmpLink) Insert(emp *Emp) {

	cur := this.Head   //辅助指针
	var pre *Emp = nil //辅助指针 pre 在 cur前面
	//如果当前的EmpLink就是一个空链表
	if cur == nil {
		this.Head = emp
		return
	}

	//插入到最前
	if emp.Id < cur.Id {
		emp.Next = cur
		this.Head = emp
		return
	}

	//如果不是一个空链表,给emp找到对应的位置并插入
	//思路：让cur和emp比较Id， pre保持在 cur前面
	for {
		if cur != nil {
			//比较
			if cur.Id >= emp.Id {
				//找到位置
				break
			}
			//节点后移，保证pre节点在cur前面
			pre = cur
			cur = cur.Next
		} else {
			break
		}
	}

	//退出时，判断是否将emp添加到链表最后

	pre.Next = emp
	emp.Next = cur

}

// 显示链表的信息
func (this *EmpLink) showLink(no int) {
	if this.Head == nil {
		fmt.Printf("链表%d 为空\n", no)
		return
	}

	//遍历当前的链表，并显示数据
	cur := this.Head //辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表%d 雇员id=%d 名字=%s -> ", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// 根据id查找对应的雇员，如果没有，返回nil
func (this *EmpLink) FindById(id int) *Emp {
	// 辅助指针
	cur := this.Head

	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 根据id查找对应的雇员，并进行删除
func (this *EmpLink) DelById(id int) {
	cur := this.Head

	if cur == nil { //空链表，直接跳出
		fmt.Println("空链表...")
		return
	}

	//只有一个节点
	if cur.Next == nil {
		if cur.Id == id {
			this.Head = nil
			return
		} else {
			fmt.Printf("链表只有一个节点，且节点id不等于%d\n", id)
			return
		}
	} else {
		//两个以上节点，且删除元素为头结点
		if cur.Id == id {
			this.Head = cur.Next
			return
		}
	}

	flag := false
	//两个以上的节点,且删除元素不是头节点
	for {
		if cur.Next == nil { //说明到链表的最后
			break
		} else if cur.Next.Id == id { //说明找到了
			flag = true
			break
		}
		cur = cur.Next
	}

	if flag {
		cur.Next = cur.Next.Next
	} else {
		fmt.Println("没有找到该节点...")
	}

}

// 定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink //具体大小按实际调整
}

// 给HashTable 编写Insert 雇员的方法
func (this *HashTable) Insert(emp *Emp) {
	//使用散列函数，确定将该雇员添加到哪个链表
	linkNo := this.HashFun(emp.Id)
	//使用对应的链表添加
	this.LinkArr[linkNo].Insert(emp)
}

// 编写一个散列方法
func (this *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对应的链表的下标
}

// 编写方法，显示hashtable的所有雇员
func (this *HashTable) ShowAll() {
	for i := 0; i < len(this.LinkArr); i++ {
		this.LinkArr[i].showLink(i)
	}
}

// 编写方法，完成查找
func (this *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	return this.LinkArr[linkNo].FindById(id)
}

// 编写方法，完成删除
func (this *HashTable) DelById(id int) (err error) {
	//使用散列函数，确定该雇员应该在哪个链表
	linkNo := this.HashFun(id)
	this.LinkArr[linkNo].DelById(id)

	return
}

func main() {

	key := ""
	id := 0
	name := ""

	var hashtable HashTable

	for {
		fmt.Println("=====雇员系统菜单======")
		fmt.Println("input  表示添加雇员")
		fmt.Println("show   表示显示雇员")
		fmt.Println("find   表示查找雇员")
		fmt.Println("delete 表示删除雇员")
		fmt.Println("exit   表示退出系统")
		fmt.Println("请输入选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
				Next: nil,
			}

			hashtable.Insert(emp)

		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				//编写方法，显示雇员信息
				emp.Info()
			}
		case "delete":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			hashtable.DelById(id)
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
