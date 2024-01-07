package main

import (
	"fmt"
)

// 累加器
func AddUpper() func(int) int {
	var n int = 10
	var str = "hello"
	return func(x int) int {
		n = n + x
		//str += "a"
		str += string(36)        // 36 => ascii => $
		fmt.Println("str=", str) //1.str="helloa" 2.str="helloaa" 3.str="helloaaa"
		return n
	}
}

func main() {
	//使用前面的代码
	f := AddUpper()
	fmt.Println(f(1)) //11
	fmt.Println(f(2)) //13
	fmt.Println(f(3)) //16

}
