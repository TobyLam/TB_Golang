package main

import (
	"fmt"
	_ "fmt" //如果我们没有使用到一个包，但是不想去掉，前面加一个下划线 _
)

func main() {
	/*//课堂练习
	var n1 int32 = 12
	var n2 int64
	var n3 int8

	n2 = int64(n1) + 20 //int32 ---> int64 错误
	n3 = int8(n1) + 20  //int32 ---> int8 错误

	fmt.Println("n2=", n2, "n3=", n3)*/

	var n1 int32 = 12
	//var n3 int8
	var n4 int8
	n4 = int8(n1) + 127 // 【编译通过，但是 结果 不是 127+12】
	//n3 = int8(n1) + 128 // 【编译不过】
	fmt.Println(n4)
}
