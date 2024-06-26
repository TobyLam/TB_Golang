package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 1) 给 Person 结构体添加 speak 方法,输出 xxx 是一个好人
func (person Person) speak() {
	fmt.Println(person.Name, "是一个好人")
}

// 2) 给 Person 结构体添加 jisuan 方法,可以计算从 1+..+1000 的结果, 说明方法体内可以函数一样，
// 进行各种运算
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 1000; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果=", res)
}

// 3) 给 Person 结构体 jisuan2 方法,该方法可以接收一个数 n，计算从 1+..+n 的结果
func (p Person) jisuan2(n int) {
	res := 0
	for i := 1; i <= n; i++ {
		res += i
	}
	fmt.Println(p.Name, "计算的结果=", res)
}

// 给Person类型绑定一个方法
func (person Person) test() {
	person.Name = "jack"
	fmt.Println("test() name=", person.Name) //输出jack
}

// 4) 给 Person 结构体添加 getSum 方法,可以计算两个数的和，并返回结果
func (p Person) getSum(n1 int, n2 int) int {
	return n1 + n2
}

type Dog struct {
}

func main() {

	var p Person
	p.Name = "tom"
	p.test() //调用方法

	fmt.Println("main() p.Name=", p.Name) //输出 tom

	//下面的使用方法都是错误的
	//test()
	//var dog Dog
	//dog.test()
	p.speak()
	p.jisuan()
	p.jisuan2(10)

	n1 := 10
	n2 := 20
	res := p.getSum(n1, n2)
	fmt.Println("res=", res)
}
