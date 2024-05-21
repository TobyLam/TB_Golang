package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 冒泡排序
func bubbleSort(arr *[11]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}

// 二分查找
func findVal(arr *[11]int, leftIndex int, rightIndex int, findValue int) {
	//中间下标
	middle := (leftIndex + rightIndex) / 2

	if leftIndex > rightIndex {
		fmt.Printf("查找%v,找不到该数\n", findValue)
		return
	}

	if (*arr)[middle] > findValue {
		findVal(arr, leftIndex, middle-1, findValue)
	} else if (*arr)[middle] < findValue {
		findVal(arr, middle+1, rightIndex, findValue)
	} else if (*arr)[middle] == findValue {
		fmt.Printf("有%v,下标为%v\n", findValue, middle)
		//数据重复时，不立即返回,遍历上下两部分，查询重复数据的下标
		//循环遍历，如果再次使用二分法，并非原数组的下标，故使用循环遍历
		for i := middle - 1; i >= 0; i-- {
			if findValue == (*arr)[i] {
				fmt.Printf("有%v,下标为%v\n", findValue, i)
			}
		}
		for i := middle + 1; i <= rightIndex; i++ {
			if findValue == (*arr)[i] {
				fmt.Printf("有%v,下标为%v\n", findValue, i)
			}
		}
	}

}

func main() {
	/**
	 * 随机生成10个整数(1-100之间)，使用冒泡排序法进行排序，然后使用二分查找法，查找是
	 * 否有90这个数，并显示下标，如果没有则提示“找不到该数”
	 */

	//声明一个数组
	arr := [11]int{}
	//生成随机数
	rand.Seed(time.Now().Unix())
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100) + 1
	}

	fmt.Println("排序前数组：", arr)
	// 冒泡排序
	bubbleSort(&arr)

	fmt.Println("排序后数组：", arr)

	findVal(&arr, 0, len(arr)-1, 99)
}
