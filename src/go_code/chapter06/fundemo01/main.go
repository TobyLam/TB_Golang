package main

import (
	"fmt"
)

// 将计算的功能，放到一个函数中，然后在需要使用，调用即可
func cal(n1 float64, n2 float64, operator byte) float64 {
	var res float64
	switch operator {
	case '+':
		res = n1 + n2
	case '-':
		res = n1 - n2
	case '*':
		res = n1 * n2
	case '/':
		res = n1 / n2
	default:
		fmt.Println("操作符号错误...")
	}

	return res
}

func main() {
	//输入两个数,再输入一个运算符(+,-,*,/)，得到结果

	var n1 float64 = 5.3
	var n2 float64 = 2.3
	var operator byte = '-'
	result := cal(n1, n2, operator)

	fmt.Println("res=", result)

	// .....
	n1 = 4.5
	n2 = 6.7
	operator = '*'
	result = cal(n1, n2, operator)

	fmt.Println("res=", result)

}
