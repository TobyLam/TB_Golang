package main

import (
	"fmt"
)

func main() {

	////1) 数组是多个相同类型数据的组合,一个数组一旦声明/定义了,其长度是固定的, 不能动态变化
	//var arr01 [3]int
	//arr01[0] = 1
	//arr01[1] = 30
	////arr01[2] = 1.1 //这里会报错,数据类型错误
	//arr01[2] = 90
	////arr01[3] = 890 //这里会报错，长度固定，不能动态变化，否则报越界
	//fmt.Println(arr01)

	//数组创建后，如果没有赋值，有默认值（零值）
	//1.数值（整数系列,浮点数系列) => 0
	//2.字符串==》 “”
	//3.布尔类型 ==> false
	var arr01 [3]float32
	var arr02 [3]string
	var arr03 [3]bool
	fmt.Printf("arr01=%v arr02=%v arr03=%v \n", arr01, arr02, arr03)
}
