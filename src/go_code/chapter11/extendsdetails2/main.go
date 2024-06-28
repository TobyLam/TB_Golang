package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	//Name string
}

func main() {
	var c C
	//如果c没有Name字段，而 A 和B 有Name,这时就必须通过指定匿名结构体名字来区分
	//所有 c.Name 就会报编译错误，这个规则对方法也适用
	c.A.Name = "tom" //error
	fmt.Println(c)
}
