package main

import (
	"fmt"
	"unsafe"
)

// 演示golang中bool类型的使用
func main() {
	var b = false
	fmt.Println("b=", b)
	//注意事项
	//1.bool类型占用存储空间是1个字节
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
	//2.bool类型的取值只能是true 或 false
	//b = 1 错误
}
