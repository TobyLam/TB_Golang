package main

import (
	"fmt"
)

/*
*
求函数值
已知 f(1)=3; f(n) = 2*f(n-1)+1;
请使用递归的思想编程，求出 f(n)的值?
*/
func f(n int) int {
	if n == 1 {
		return 3
	} else {
		return 2*f(n-1) + 1
	}
}

func main() {
	var i int = 3
	fmt.Printf("当i等于%v时，f(n)=%v", i, f(i))
	fmt.Println("")
	i = 5
	fmt.Printf("当i等于%v时，f(n)=%v", i, f(i))
	fmt.Println("")
}
