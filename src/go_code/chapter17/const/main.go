package main

import (
	"fmt"
)

func main() {
	var num int
	num = 90 //ok
	//常量声明的时候，必须赋值。
	const tax int = 0
	//常量不能修改
	//tax = 10 //error
	fmt.Println(num, tax)

	//常量只能修饰bool、数值类型(int,float系列)、string类型
	const a = 9 //ok
	//const b = num/3 //error
	const b = a / 3
	fmt.Println(b)

	const (
		a1 = iota
		b1
		c1
		d1
		e1, f1 = iota, iota
	)
	fmt.Println(a1, b1, c1, d1, e1, f1)
}
