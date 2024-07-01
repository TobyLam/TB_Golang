package main

import (
	"fmt"
)

type Student struct {
}

// 编写一个函数，可以判断输入的参数是什么类型
func TypeJudge(items ...interface{}) {
	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index+1, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index+1, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index+1, x)
		case int, int8, int16, int32, int64:
			fmt.Printf("第%v个参数是int类型，值是%v\n", index+1, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index+1, x)
		case Student:
			fmt.Printf("第%v个参数是Student类型，值是%v\n", index+1, x)
		case *Student:
			fmt.Printf("第%v个参数是*Student类型，值是%v\n", index+1, x)
		default:
			fmt.Printf("第%v个参数是未知类型\n", index+1)
		}
	}
}

func main() {
	var n1 float32 = 1.1
	var n2 float64 = 2.2
	var n3 int8 = 3
	var n4 int16 = 4
	var n5 int32 = 5
	var n6 int64 = 6
	var n7 int = 7
	var n8 string = "8"
	address := "北京"
	n9 := 200

	stu1 := Student{}
	stu2 := &Student{}

	TypeJudge(n1, n2, n3, n4, n5, n6, n7, n8, address, n9, stu1, stu2)

}
