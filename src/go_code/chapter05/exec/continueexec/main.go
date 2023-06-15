package main

import (
	"fmt"
)

func main() {
	//continue 实现 打印 1——100 之内的奇数[要求使用 for 循环+continue]
	/*for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("奇数是", i)
	}*/

	//从键盘读入个数不确定的整数，并判断读入的正数和负数的个数，输入为 0 时结束程序
	var positiveCount int
	var nagativeCount int
	var num int
	for {
		fmt.Println("输入一个整数，输入0结束")
		fmt.Scanln(&num)
		if num == 0 {
			break //终止for循环
		}

		if num > 0 {
			positiveCount++
			continue //大于0后没必要继续判断小于0
		}

		nagativeCount++

	}
	fmt.Printf("正数的个数是%v,负数的个数是%v \n", positiveCount, nagativeCount)

	//某人有 100,000 元,每经过一次路口，需要交费,规则如下:
	//当现金>50000 时,每次交 5%
	//当现金<=50000 时,每次交 1000
	//编程计算该人可以经过多少次路口,使用 for break 方式完成

	var cash float64 = 100000.0
	count := 0
	for {
		if cash <= 0 {
			break
		}
		if cash > 50000 {
			cash -= cash * 0.05
			count++
			continue
		}

		cash -= 1000
		count++
	}
	fmt.Printf("该人可以经过%v次路口", count)
}
