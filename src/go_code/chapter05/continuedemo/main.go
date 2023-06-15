package main

import (
	"fmt"
)

func main() {

	//continue案例
	//指定标签的形式使用 continue
	//lable2:
	for i := 0; i < 4; i++ {
		//lable1: //设置一个标签
		for j := 0; j < 10; j++ {
			if j == 2 {
				//break lable1 // break 默认跳出最近的for循环
				continue
			}
			fmt.Println("j=", j)
		}
	}
}
