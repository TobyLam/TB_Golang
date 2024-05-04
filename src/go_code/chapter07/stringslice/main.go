package main

import (
	"fmt"
)

func main() {
	//string底层是一个byte数组，因此string也可以进行切片处理
	str := "hello@TobyLam"

	//使用切片获取到TobyLam
	slice := str[6:]

	fmt.Println("slice=", slice)

	var d[];

}
