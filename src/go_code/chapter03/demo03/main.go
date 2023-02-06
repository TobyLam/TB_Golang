package main

import "fmt"

// 定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

// 上面的声明方式，也可以改成一次性声明
var (
	n3    = 100
	n4    = 200
	name2 = "mary"
)

func main() {
	//该案例演示golang如何一次性声明多个变量
	//var n1, n2, n3 int
	//fmt.Println("n1=", n1, "n2=", n2, "n3=", n3)

	//一次性声明多个变量的方法2
	//var n1, name, n3 = 100, "tom", 888
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//一次性声明多个变量的方法3,同样可以使用类型推导
	//n1, name, n3 := 100, "tom~", 888
	//fmt.Println("n1=", n1, "name=", name, "n3=", n3)

	//输出全局变量
	fmt.Println("n1=", n1, "n2=", n2, "name=", name)
	fmt.Println("n3=", n3, "n4=", n4, "name2=", name2)

}
