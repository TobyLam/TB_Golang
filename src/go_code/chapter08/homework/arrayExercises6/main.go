package main

import (
	"fmt"
)

func findAA(arr [10]string, findVal string) {

	result := 0
	for i, v := range arr {
		if v == findVal {
			result = 1
			fmt.Printf("%v对应的下标为%v\n", findVal, i)
		}
	}
	if result == 0 {
		fmt.Printf("%v没找到\n", findVal)
	}

}

func main() {
	/**
	    * 试写出实现查找的核心代码，比如已知数组arr[10]string;里面保存了十个元素，现要查找
		" "AA"在其中是否存在，打印提示，如果有多个"AA”,也要找到对应的下标。
	*/
	arr := [10]string{"AA", "AB", "BB", "BBD", "AA", "G", "UI", "AA", "AA", "HELLO"}
	findAA(arr, "AA")
	findAA(arr, "AB")
}
