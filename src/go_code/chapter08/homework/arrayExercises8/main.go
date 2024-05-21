package main

import (
	"fmt"
)

func findMaxMin(arr [5]float64) {
	min := arr[0]
	minIndex := 0
	max := arr[0]
	maxIndex := 0

	for i, v := range arr {

		if v >= max {
			max = v
			maxIndex = i
		}
		if v <= min {
			min = v
			minIndex = i
		}
	}
	fmt.Printf("最大的数是%v,下标为%v,最小的数是%v,下标为%v\n", max, maxIndex, min, minIndex)
}

func main() {

	/**
	 * 编写一个函数，可以接收一个数组，该数组有5个数，请找出最大的数和最小的数和对应的数组下标是多少?
	 */

	arr := [5]float64{}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("请输入第%v个数的值:\n", i+1)
		fmt.Scanln(&arr[i])
	}

	fmt.Println("原数组=", arr)

	findMaxMin(arr)

}
