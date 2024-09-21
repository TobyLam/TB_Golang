package main

import (
	"fmt"
	"sort"
)

var name = "tom" //全局变量

// Name := "mary"  //这个地方会报错,执行语句不能放在函数外

func test01() {
	fmt.Println(name) //tom
}

func test02() {
	name := "jack"
	fmt.Println(name) // jack
}

func main() {
	fmt.Println(name) //tom
	test01()          //tom
	test02()          //jack
	test01()          //tom

	mySlice := []int{1, 0, -123}
	sort.Ints(mySlice)
}
