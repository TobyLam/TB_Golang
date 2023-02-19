package main

import "fmt"

// 演示golang中字符类型使用
func main() {

	var c1 byte = 'a'
	var c2 byte = '0' //字符的0

	//当我们直接输出byte值，就是输出了的对应的字符的码值
	fmt.Println("c1=", c1, "c2=", c2)

	//如果我们希望输出对应字符，需要使用格式化输出
	fmt.Printf("c1=%c c2=%c \n", c1, c2)

	//var c3 byte = '北' //overflow溢出
	var c3 int = '北'
	fmt.Printf("c3=%c c3对应码值=%d \n", c3, c3)

	//可以直接给某个变量赋一个数字， 然后按格式化输出时%c， 会输出该数字对应的 unicode 字符
	var c4 int = 120 // 22269 -> '国' 120->'x'
	fmt.Printf("c4=%c\n", c4)

	//字符类型是可以进行运算的， 相当于一个整数， 因为它都对应有 Unicode 码
	var n1 = 10 + 'a' // 10 +97
	fmt.Println("n1=", n1)
	fmt.Printf("n1=%c", n1)
}
