package main

import (
	"fmt"
)

// 一个函数 test
func test(n1 int) {
	n1 = n1 + 1
	fmt.Println("test() n1=", n1)
}

// 一个函数 getSum
func getSum(n1 int, n2 int) int {
	sum := n1 + n2
	fmt.Println("getSum sum = ", sum)
	//当函数有return 语句时，就是将结果返回给调用者
	//即谁调用我，就返回给谁
	return sum
}

// 请编写要给函数，可以计算两个数的和和差，并返回结果。
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}

func main() {

	n1 := 10
	//调用test
	test(n1)

	fmt.Println("main() n1=", n1)

	sum := getSum(1, 2)
	fmt.Println("main sum = ", sum)

	//调用getSumAndSub
	res1, res2 := getSumAndSub(1, 2)
	fmt.Printf("res1=%v res2=%v\n", res1, res2)

	//希望忽略某个返回值，则使用 _ 符号表示占位忽略
	_, res3 := getSumAndSub(1, 2)
	fmt.Printf("res3=%v\n", res3)

	_, _ = getSumAndSub(3, 4) // 等价于  getSumAndSub(3,4)
}
