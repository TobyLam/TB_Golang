package main

import (
	"fmt"
)

func main() {

	var intArr [3]int //int占8个字节
	//当定义完数组后，其实数组的各个元素有默认值 0
	fmt.Println(intArr)

	intArr[0] = 10
	intArr[1] = 20
	intArr[2] = 30
	fmt.Println(intArr)
	fmt.Printf("intArr的地址=%p intArr[0] 的地址%p intArr[1] 地址%p intArr[2]的地址%p\n", &intArr, &intArr[0], &intArr[1], &intArr[2])

	//从终端循环输入 5 个成绩，保存到 float64 数组,并输出
	//var score [5]float64
	//
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("请输入第%d个元素的值", i+1)
	//	fmt.Scanln(&score[i])
	//}
	//
	////遍历数组打印
	//for i := 0; i < len(score); i++ {
	//	fmt.Printf("score[%d]=%v", i, score[i])
	//}

	//四种初始化数组的方式
	var numArr01 [3]int = [3]int{1, 2, 3}
	fmt.Println("numArr01=", numArr01)

	var numArr02 = [3]int{11, 22, 33}
	fmt.Println("numArr02=", numArr02)

	//这里的 [...]是规定的写法
	var numArr03 = [...]int{111, 222, 333}
	fmt.Println("numArr03=", numArr03)

	var numArr04 = [...]int{1: 1111, 0: 222, 2: 333}
	fmt.Println("numArr04=", numArr04)

	strArr05 := [...]string{1: "tom", 0: "jack", 2: "mary"}
	fmt.Println("numArr05=", strArr05)
}
