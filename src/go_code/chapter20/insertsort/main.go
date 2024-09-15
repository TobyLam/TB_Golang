package main

import (
	"fmt"
	"math/rand"
	"time"
)

/**
把 n 个待排序的元素看成为一个有序表和一个无序表，
开始时有序表中只包含一个元素，无序表中包含有 n-1 个元素，排序过程中每次从无序表中取出第一个
元素，把它的排序码依次与有序表元素的排序码进行比较，将它插入到有序表中的适当位置，使之成
为新的有序表。
*/

// 插入排序【升序】
func InsertSortAsc(arr *[7]int) {

	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1

		for insertIndex >= 0 && arr[insertIndex] > insertVal {
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}

		arr[insertIndex+1] = insertVal

		fmt.Printf("第%d次插入后 %v\n", i, *arr)
	}
}

// 插入排序【降序】
func InsertSortDesc(arr *[80000]int) {

	//完成第一次,给第二个元素找到合适的位置并插入,默认第一个元素即为有序列表
	for i := 1; i < len(arr); i++ {

		insertVal := arr[i]  //排序值
		insertIndex := i - 1 //开始比较的下标，排序值前一个位置的数据开始比较

		//从大到小（do...while...)

		//从后往前比较
		//insertIndex 未超出数组（>=0） 且比较位置的数据 小于 比较的数据
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] //数据后移
			insertIndex--                         //比较位前移
		}

		//插入
		if insertIndex+1 != i { //插入数据就是最后一个，不需要修改
			arr[insertIndex+1] = insertVal
		}

		//fmt.Printf("第%d次插入后 %v\n", i, *arr)
	}

}

func main() {
	//arr := [7]int{23, 0, 12, 56, 34, -1, 55}

	//fmt.Println("原始数组=", arr)

	var arr [80000]int
	for i := 0; i < 80000; i++ {
		arr[i] = rand.Intn(900000)
	}

	start := time.Now().Unix()
	InsertSortDesc(&arr)
	end := time.Now().Unix()
	fmt.Printf("插入排序耗时=%d秒\n", end-start)

	fmt.Println("main 函数")
	//fmt.Println(arr)

	//arr2 := [7]int{23, 0, 12, 56, 34, -1, 55}
	//
	//fmt.Println("原始数组=", arr2)
	//InsertSortAsc(&arr2)
	//
	//fmt.Println("main 函数")
	//fmt.Println(arr2)

}
