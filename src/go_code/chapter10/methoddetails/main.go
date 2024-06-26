package main

import (
	"fmt"
)

/*
*3) Golang 中的方法作用在指定的数据类型上的(即：和指定的数据类型绑定)，因此自定义类型，
都可以有方法，而不仅仅是 struct， 比如 int , float32 等都可以有方法
*/
type integer int

func (i integer) print() {
	fmt.Println("i=", i)
}

type Student struct {
	Name string
	Age  int
}

// 给Student实现方法String()
func (stu *Student) String() string {
	str := fmt.Sprintf("Name=[%v] Age=[%v]", stu.Name, stu.Age)
	return str
}

// 编写一个方法，可以改变id的值
func (i *integer) change() {
	*i = *i + 1
}
func main() {
	var i integer = 10
	i.print()
	i.change()
	fmt.Println("i=", i)

	//定义一个Student变量
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	//如果你实现了 *Student 类型的 String方法，就会自动调用
	fmt.Println(&stu)
}
