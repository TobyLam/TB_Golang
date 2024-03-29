package main

import (
	"fmt"
)

// 在 Go 中，函数也是一种数据类型，
// 可以赋值给一个变量，则该变量就是一个函数类型的变量了。通过该变量可以对函数调用

func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 自定义数据类型
// 这时 myFun 就是 func(int,int) int类型
type myFunType func(int, int) int

// 函数既然是一种数据类型，因此在 Go 中，函数可以作为形参，并且调用
func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 支持对函数返回值命名
// 不需要再关心 ruturn的顺序错误
func getSumAndSub(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

// 可变参数
// 编写一个函数sum,可以求出 1到多个int的和
func sum(n1 int, args ...int) int {
	sum := n1
	//遍历args
	for i := 0; i < len(args); i++ {
		sum += args[i] //args[0] 表示取出args切片的第一个元素值，其他依此类推
	}
	return sum
}

func main() {
	a := getSum
	fmt.Printf("a的类型是%T,\ngetSum类型是%T\n", a, getSum)

	res := a(10, 40) //等价 res := getSum(10,40)
	fmt.Println("res=", res)

	//看案例
	res2 := myFun(getSum, 50, 60)
	fmt.Println("res2=", res2)

	// 给int取了别名  在go中 myInt h和 int 虽然都是int类型，但是go认为myInt和int两个类型
	type myInt int

	var num1 myInt //
	var num2 int

	num1 = 40
	num2 = int(num1) //注意这里依然需要显式转换，go认为myInt和int两个类型
	fmt.Println("num1=", num1, "num2=", num2)

	//看案例
	res3 := myFun2(getSum, 500, 600)
	fmt.Println("res3=", res3)

	//
	a1, b1 := getSumAndSub(1, 2)
	fmt.Printf("a1=%v b1=%v\n", a1, b1)

	//测试一下可变参数的使用
	res4 := sum(10, 0, -1, 90, 10)
	fmt.Println("res4=", res4)

}
