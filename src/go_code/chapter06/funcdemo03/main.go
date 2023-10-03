package main

import "fmt"

func test(n int) {
	if n > 2 {
		n--
		test(n)
	}
	fmt.Println("n=", n)
}

func test2(n int) {
	if n > 2 {
		n--
		test2(n)
	} else {
		fmt.Println("n=", n)
	}
}

/**
 * 递归调用demo
 */
func main() {

	//看一段代码
	test(4)  // 输出 n=2 n=2 n=3
	test2(4) // 输出 n=2
}
