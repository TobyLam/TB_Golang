package main

import (
	"fmt"
)

func main() {
	//判断一个三位数是否是水仙花数
	var x int
	var a, b, c int

	fmt.Println("请输入一个三位数")
	fmt.Scanln(&x)
	a = x / 100
	b = x % 100 / 10
	c = x % 10

	//fmt.Printf("x=%v,a=%v,b=%v,c=%v", x, a, b, c)
	if a*a*a+b*b*b+c*c*c == x {
		fmt.Printf("%v是水仙花数", x)
	} else {
		fmt.Printf("%v不是水仙花数", x)
	}
}
