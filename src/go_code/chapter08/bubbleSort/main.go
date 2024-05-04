package main

import (
	"fmt"
)

// 冒泡排序
func BubbleSort(arr *[5]int) {
	fmt.Println("排序前arr=", (*arr))

	//临时变量写法
	//temp := 0 //临时变量（用于交换）
	////完成第一轮排序（外层1次)
	//for j := 0; j < 4; j++ {
	//	if (*arr)[j] > (*arr)[j+1] {
	//		//交换
	//		temp = (*arr)[j+1]
	//		(*arr)[j+1] = (*arr)[j]
	//		(*arr)[j] = temp
	//	}
	//}

	//非临时变量写法
	for i := 0; i < len(*arr)-1; i++ {
		//完成第i轮排序（外层i次)
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
	fmt.Println("排序后arr=", (*arr))
}

func main() {
	//定义数组
	arr := [5]int{24, 69, 80, 57, 13}
	//将数组传递给一个函数，完成排序

	BubbleSort(&arr)

	fmt.Println("main arr=", arr) //有序？ 是有序的
}
