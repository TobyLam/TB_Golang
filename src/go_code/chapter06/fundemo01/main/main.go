package main

import (
	"fmt"
	// 别名 util ,原包名不能再使用
	util "go_code/chapter06/fundemo01/utils"
)

func main() {

	fmt.Println("utils.go Num=", util.Num1)
	//输入两个数,再输入一个运算符(+,-,*,/)，得到结果

	var n1 float64 = 5.3
	var n2 float64 = 2.3
	var operator byte = '-'
	result := util.Cal(n1, n2, operator)

	fmt.Println("res=~", result)

	// .....
	n1 = 4.5
	n2 = 6.7
	operator = '*'
	result = util.Cal(n1, n2, operator)

	fmt.Printf("res=~%.2f", result)

	//util.SayOk()

}
