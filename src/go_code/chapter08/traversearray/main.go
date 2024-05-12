package main

import (
	"fmt"
)

func main() {

	//二维数组的遍历

	var arr = [...][3]int{{1, 2, 3}, {4, 5, 6}}

	//for循环遍历
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%v\t", arr[i][j])
		}
		fmt.Println()
	}

	//for-range遍历二维数组
	for i, v := range arr {
		for j, v2 := range v {
			fmt.Printf("arr[%v][%v]=%v \t", i, j, v2)
		}
		fmt.Println()
	}
}
