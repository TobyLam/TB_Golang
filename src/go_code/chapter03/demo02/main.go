package main

import "fmt"

func main() {
	//golang的变量使用方式
	//第一种：指定变量类型，声明后若不赋值，使用默认值
	// int 的默认值是0
	var i int
	fmt.Println("i=", i)

	//第二种：根据值自行判定变量类型（类型推导）
	var num = 10.11
	fmt.Println("num=", num)

	//第三种： 省略 var, 注意 :=左侧的变量不应该是已经声明过的， 否则会导致编译错误
	//下面方式等价于 var name string   name = "tom"
	// := 的 :不能省略 否则就是错误的
	name := "tom"
	fmt.Println("name=", name)
}
