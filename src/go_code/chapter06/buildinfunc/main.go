package main

import (
	"fmt"
)

func main() {

	num1 := 100
	fmt.Printf("num1的类型 %T, num1的值=%v, num1的地址%v \n", num1, num1, &num1)

	num2 := new(int) // *int
	// num2 的类型%T => *int
	// num2 的值 = 地址？
	// num2 的地址%v = 地址
	// num2 指向的值 = 0
	*num2 = 100
	fmt.Printf("num2的类型 %T, num2的值=%v, num2的地址%v \nnum2这个指针，指向的值=%v\n",
		num2, num2, &num2, *num2)
}
