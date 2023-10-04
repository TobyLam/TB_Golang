package main

import (
	"fmt"
)

/*
*
题 3：猴子吃桃子问题
有一堆桃子，猴子第一天吃了其中的一半，并再多吃了一个！以后每天猴子都吃其中的一半，然后
再多吃一个。当到第十天时，想再吃时（还没吃），发现只有 1 个桃子了。问题：最初共多少个桃子？
*/
func monkeyEatPeach(n int) int {

	/**
	1) 第 10 天只有一个桃子
	2) 第 9 天有几个桃子 = (第 10 天桃子数量 + 1) * 2
	3) 规律: 第 n 天的桃子数据 peach(n) = (peach(n+1) + 1) * 2
	*/
	if n > 10 || n < 1 {
		fmt.Println("输入的天数不对")
		return 0
	}

	if n == 10 {
		return 1
	} else {
		return 2*monkeyEatPeach(n+1) + 2
	}

}

func main() {
	fmt.Printf("最初共有%v个桃子", monkeyEatPeach(1))
}
