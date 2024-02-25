package main

import (
	"fmt"
)

func main() {

	//for-fange遍历数组
	heroes := [...]string{"孔明", "子龙", "孟起"}
	fmt.Println(heroes)
	//使用常规的方式遍历... len

	for i, v := range heroes {
		fmt.Printf("i=%v v=%v\n", i, v)
		fmt.Printf("heroes[%d]=%v\n", i, heroes[i])
	}

	for _, v := range heroes {
		fmt.Printf("元素的值=%v\n", v)

	}

}
