package main

import "fmt"

// 编写一个函数调用九九乘法表
func printMulti(num int) {
	// 打印出九九乘法表
	fmt.Println("打印九九乘法表")
	for s := 1; s <= num; s++ {
		for e := 1; e <= s; e++ {
			fmt.Printf("%v * %v = %v \t", s, e, s*e)
		}
		fmt.Println()
	}
}
func main() {
	//从终端输入一个整数表示要打印的乘法表对应的数
	var num int
	fmt.Println("请输入XX乘法表的对应数")
	fmt.Scanln(&num)
	printMulti(num)
}
