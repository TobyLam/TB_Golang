package main

import (
	"errors"
	"fmt"
)

// 使用数组模拟一个栈的使用
type Stack struct {
	maxTop int    //表示栈最大的可存放个数
	Top    int    //表示栈顶,因为栈底固定，直接使用Top
	arr    [5]int //数组模拟栈
}

/**
 * 入栈
 */
func (this *Stack) Push(val int) (err error) {

	//判断栈是否已满
	if this.IsFull() {
		fmt.Println("stack full")
		return errors.New("stack full")
	}

	this.Top++
	//放入数据
	this.arr[this.Top] = val

	return
}

/**
 * 出栈
 */
func (this *Stack) Pop() (val int, err error) {

	//判断是否为空
	if this.IsEmpty() {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}

	//先取值，再 this.Top--
	val = this.arr[this.Top]
	this.Top--

	return val, nil
}

// 判断栈是否已满
func (this *Stack) IsFull() bool {
	return this.Top == this.maxTop-1
}

// 判断栈是否为空
func (this *Stack) IsEmpty() bool {
	return this.Top == -1
}

// 遍历栈，从栈顶开始遍历
func (this *Stack) List() {
	//先判断栈是否为空
	if this.IsEmpty() {
		fmt.Println("stack empty")
		return
	}
	fmt.Println("栈的情况如下：")
	for i := this.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, this.arr[i])
	}
}

func main() {

	stack := &Stack{
		maxTop: 5,  //最多存放5个数到栈中
		Top:    -1, //当栈顶为-1时，表示栈为空
	}

	//入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)

	//显示
	stack.List()

	val, _ := stack.Pop()
	fmt.Println("出栈val=", val) //先入后出，5

	//显示
	stack.List() //4,3,2,1

	val, _ = stack.Pop()
	fmt.Println("出栈val=", val) //先入后出，4

	//显示
	stack.List() //3,2,1

	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	val, _ = stack.Pop()
	fmt.Println("出栈val=", val) //先入后出，4

	//显示
	stack.List() //3,2,1
}
