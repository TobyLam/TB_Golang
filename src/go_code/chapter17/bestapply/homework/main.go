package main

import (
	"reflect"
)

// 1）编写一个结构体，有两个字段 Num1,和Num2
type Cal struct {
	Num1 int
	Num2 int
}

// 2)方法GetSub(name string)
func (c Cal) GetSub(name string) {

}

// 3)使用反射遍历Cal结构体所有的字段信息

// 4)使用反射机制完成对GetSub的调用，输出形式为
// "tom 完成了减法运行，8-3=5

func main() {

}
