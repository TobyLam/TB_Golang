package utils

import "fmt"

var Num1 int = 300

// 将计算的功能，放到一个函数中，然后在需要使用，调用即可
// 为了让其它包的文件使用Cal函数，需要将C大写，类似其它语言的public [专业术语: 该函数可导出]
func Cal(n1 float64, n2 float64, operator byte) float64 {
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
