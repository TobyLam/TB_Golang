package main

import (
	"fmt"
)

func main() {
	//试保存1 3 5 7 9 五个奇数到数组，并倒序打印
	arr := [5]int{}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("请输入第%v个奇数\n", i+1)
		fmt.Scanln(&arr[i])
	}

	for i := len(arr) - 1; i >= 0; i-- {
		fmt.Println(arr[i])
	}

	//排序方式
	// for i := 0; i < len(arr); i++ {
	// 	for j := 0; j < len(arr)-i-1; j++ {
	// 		if arr[j] < arr[j+1] {
	// 			arr[j], arr[j+1] = arr[j+1], arr[j]
	// 		}
	// 	}
	// }

	//fmt.Println("倒序打印arr=", arr)
}
