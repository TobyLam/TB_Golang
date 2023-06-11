package main

import (
	"fmt"
)

func main() {
	//使用 for 循环完成下面的案例请编写一个程序，可以接收一个整数,表示层数，打印出金字塔

	//思路
	//1. 打印一个矩形
	/*
	 ***    | *     1个*  |    *    1个*：2i-1  空格： 2（总层数-当前层数）
	 ***    | **    2个*  |   ***   3个*：2i-1  空格： 1
	 ***    | ***   3个*  |  *****  5个*：2i-1  空格： 0
	 */
	//2. 打印半个金字塔
	//3. 打印整个金字塔
	//4. 将层数做成变量

	//5. 打印空心金字塔
	// 分析：给每行打印*号时，需要考虑是打印*还是打印空格
	// 每行第一个和最后一个打印*，其余空格，最后一层除外
	/*var totalLevel int = 0
	fmt.Println("请输入金字塔层数: \n")
	fmt.Scanln(&totalLevel)*/
	var totalLevel int = 5
	fmt.Println("打印空心金字塔")
	// i表示层数
	for i := 1; i <= totalLevel; i++ {
		// 打印*前先打印空格
		for k := 1; k <= totalLevel-i; k++ {
			fmt.Print(" ")
		}
		// j表示每层打印多少个*
		for j := 1; j <= 2*i-1; j++ {
			if j == 1 || j == 2*i-1 || i == totalLevel {
				if i == totalLevel {
					if j%2 == 0 {
						fmt.Print(" ")
					} else {
						fmt.Print("*")
					}
				} else {
					fmt.Print("*")
				}

			} else {
				fmt.Print(" ")
			}

		}
		fmt.Println("")
	}
	fmt.Println()
	// 打印空心菱形
	// 知道最大一行的*数量

	/**
		上半部分最大行数 num

				行数           空格数                  *数
		上半
		  *      1      n       2     num-n          1     2n - 1
		 ***     2      n       1     num-n          3     2n - 1
		*****    3      n       0     num-n          5     2n - 1
	    下半 [倒三角形] 剩余行数 num - 1, 循环次数= num -1
		 ***     1      n       1      n             3     2num - 2n - 1
		  *      2      n       2      n             1     2num - 2n - 1
	*/

	/*var num int
	fmt.Println("请输入菱形上半部分行数：")
	fmt.Scanln(&num)*/
	var num int = 8
	fmt.Println("打印空心菱形")
	//上半部分
	for m := 1; m <= num; m++ {
		for k := 1; k <= num-m; k++ {
			fmt.Print(" ") //打印空格
		}
		for n := 1; n <= 2*m-1; n++ {
			if n == 1 || n == 2*m-1 { //打印边缘星号
				fmt.Print("*")
			} else { //打印星号里面的空格
				fmt.Print(" ")
				//fmt.Print("*")
			}
		}
		fmt.Println()
	}
	//下半部分
	for m := num - 1; m >= 1; m-- { //下半部分层数比上半部分少一层
		for k := 1; k <= num-m; k++ {
			fmt.Print(" ")
		}
		for n := 1; n <= 2*m-1; n++ {
			if n == 1 || n == 2*m-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
				//fmt.Print("*")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	// 打印出九九乘法表
	fmt.Println("打印九九乘法表")
	for s := 1; s <= 9; s++ {
		for e := 1; e <= s; e++ {
			fmt.Printf("%v * %v = %v \t", s, e, s*e)
		}
		fmt.Println()
	}

}
