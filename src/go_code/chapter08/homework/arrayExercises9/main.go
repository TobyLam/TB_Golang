package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/**
	 * 定义一个数组，并给出8个整数，求该数组中大于平均值的数的个数，和小于平均值的数的个数。
	 */
	arr := [8]int{}
	rand.Seed(time.Now().Unix())
	sum := 0
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
		sum += arr[i]
	}
	avg := float64(sum) / float64(len(arr))
	fmt.Println("原数组=", arr)
	fmt.Printf("总和=%v,平均值=%v\n", sum, avg)

	ltAvgNum := 0
	gtAvgNum := 0
	for _, v := range arr {
		if float64(v) > avg {
			gtAvgNum++
		}
		if float64(v) < avg {
			ltAvgNum++
		}
	}

	fmt.Printf("该数组中大于平均值的数的个数是%v,小于平均值的数的个数是%v\n", gtAvgNum, ltAvgNum)
}
