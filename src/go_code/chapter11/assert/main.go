package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point //ok

	//如果将 a 赋给一个Point变量？
	var b Point
	// b = a // ok ?  =>error
	b = a.(Point) //ok
	fmt.Println(b)

	// //类型断言的其他案例
	// var x interface{}
	// var b2 float32 = 1.1
	// x = b2 // 空接口可以接收任意类型
	// // x=>float32 [使用类型断言]
	// y := x.(float64)
	// fmt.Printf("y的类型是%T 值是=%v\n", y, y)

	//类型断言(带检测的)
	var x interface{}
	var b2 float32 = 2.1
	x = b2 // 空接口可以接收任意类型
	// x=>float32 [使用类型断言]
	if y, ok := x.(float32); ok {
		fmt.Println("convert success")
		fmt.Printf("y的类型是%T 值是=%v\n", y, y)
	} else {
		fmt.Println("convert fail")
	}

	fmt.Println("继续执行...")

}
