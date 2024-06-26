package main

import (
	"fmt"
)

type Circle struct {
	radius float64
}

// 2) 声明一个方法 area 和 Circle 绑定，可以返回面积。
func (c Circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

// 为了提高效率，通常我们方法和结构体的指针类型绑定
func (c *Circle) area2() float64 {
	//因为c是指针，因此标准的访问其字段的方式是（*c).redius
	//return 3.14 * (*c).radius * (*c).radius
	// (*c).radius 等价 c.radius
	fmt.Printf("c 是 *Circle 指向的地址=%p\n", c)
	c.radius = 10
	return 3.14 * c.radius * c.radius

}

func main() {

	//1) 声明一个结构体 Circle, 字段为 radius
	//2) 声明一个方法 area 和 Circle 绑定，可以返回面积。
	//3) 提示：画出 area 执行过程+说明

	//创建一个Circle 变量
	var c Circle
	fmt.Printf("main c结构体变量地址%p\n", &c)
	c.radius = 5.0
	res := c.area()
	//res2 := (&c).area2()
	//编译器底层做了优化，（&c）.area2() 等价于 c.area()
	//因为编译器会自动的给加上&c
	res2 := c.area2()
	fmt.Println("面积是=", res)
	fmt.Println("面积2=", res2)
	fmt.Println("c.radius=", c.radius)
}
