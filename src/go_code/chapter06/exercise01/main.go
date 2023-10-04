package main

import (
	"fmt"
)

/*
  - 题 1：斐波那契数
    请使用递归的方式，求出斐波那契数 1,1,2,3,5,8,13... 给你一个整数 n，求出它的斐波那契数是多少？
*/
func fbn(n int) int {
	/**
	思路:
	1) 当 n == 1 || n ==2 , 返回 1
	2) 当 n >= 2, 返回 前面两个数的和 f(n-1) + f(n-2)
	*/
	if n == 1 || n == 2 {
		return 1
	} else {
		return fbn(n-1) + fbn(n-2)
	}
}

func main() {
	//测试
	n := 10
	for i := 1; i <= n; i++ {
		fmt.Printf("斐波那契数列第%v个数的值为%v", i, fbn(i))
		fmt.Println("")
	}

}
