package main

import (
	"fmt"
)

// 函数
func test01(arr [3]int) { // [3]int [4]int 数组长度是数组类型的一部分，例子中的3、4类型不一样
	arr[0] = 88
}

// 函数
func test02(arr *[3]int) {
	fmt.Printf("arr指针的地址=%p\n", &arr)
	(*arr)[0] = 88 //11
}

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

	//数组的下标是从0开始的

	//var arr04 [3]string //0-2
	//var index int = 3
	//arr04[index] = "tom" // 因为下标是0-2，因此arr04[3]就越界

	//arr := [3]int{11, 22, 33}
	//test01(arr)
	//fmt.Println("main arr=", arr) //

	arr := [3]int{11, 22, 33}
	fmt.Printf("arr 的地址=%p\n", &arr)
	test02(&arr)
	fmt.Println("main arr=", arr) //
}
