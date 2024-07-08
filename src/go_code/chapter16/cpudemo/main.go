package main

import (
	"fmt"
	"runtime"
)

func main() {

	cpuNum := runtime.NumCPU()
	fmt.Println("cpuNum=", cpuNum)

	//可以设置使用多少个cup
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}
