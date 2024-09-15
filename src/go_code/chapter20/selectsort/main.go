package main

import (
	"fmt"
	"math/rand"
	"time"
)

//选择排序，遍历选择最值，然后和最值交换位置
//然后从剩余位置继续选择最值，和最值交换位置
//以此类推...

//效率比较： 选择排序在获取到最值之后才进行交换，冒泡排序比较完成则交换，选择排序的效率更高

// 编写函数selectSort
func selectSort(arr *[80000]int) {

	//标准的访问方式
	//(*arr)[1] = 600
	//arr[1] = 600

	for j := 0; j < len(arr)-1; j++ {

		//1.假设当前的值为最值，当前的索引为最值的索引
		max := arr[j]
		maxIndex := j

		//2.遍历后面元素 （j+1）----[len(arr)-1] 比较
		for i := j + 1; i < len(arr); i++ {
			if max < arr[i] { //找到真正的最大值
				max = arr[i]
				maxIndex = i
			}
		}

		//3.交换
		if maxIndex != j {
			arr[j], arr[maxIndex] = arr[maxIndex], arr[j]
		}

		//fmt.Printf("第%d次 %v\n  ", j+1, *arr)
	}

}

func main() {
	//定义一个数组，从大到小排
	//arr := [6]int{10, 34, 19, 100, 80, 789}

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	selectSort(&arr)
	end := time.Now().Unix()
	fmt.Printf("选择排序耗时=%d秒\n", end-start)

	//遍历
	//fmt.Println(arr)
}
