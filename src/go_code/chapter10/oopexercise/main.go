package main

import (
	"fmt"
)

/*
*学生案例
1) 编写一个 Student 结构体，包含 name、gender、age、id、score 字段，分别为 string、string、int、
int、float64 类型。
2) 结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
3) 在 main 方法中，创建 Student 结构体实例(变量)，并访问 say 方法，并将调用结果打印输出。
*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (student *Student) say() string {

	infoStr := fmt.Sprintf("student的信息 name=[%v] gender=[%v] age=[%v] id=[%v] score=[%v]", student.name, student.gender,
		student.age, student.id, student.score)
	return infoStr
}

/**
小狗案例
*/
// 1) 编写一个 Dog 结构体，包含 name、age、weight 字段
// 2) 结构体中声明一个 say 方法，返回 string 类型，方法返回信息中包含所有字段值。
// 3) 在 main 方法中，创建 Dog 结构体实例(变量)，并访问 say 方法，将调用结果打印输出。
type Dog struct {
	name   string
	age    int
	weight float64
}

func (dog *Dog) dogSay() string {
	infoStr := fmt.Sprintf("dog的信息 name=[%v] age=[%v] weight=[%v]", dog.name, dog.age, dog.weight)
	return infoStr
}

/*
*
盒子案例
1) 编程创建一个 Box 结构体，在其中声明三个字段表示一个立方体的长、宽和高，长宽高要从终
端获取
2) 声明一个方法获取立方体的体积。
3) 创建一个 Box 结构体变量，打印给定尺寸的立方体的体积
*/
type Box struct {
	len    float64
	width  float64
	height float64
}

func (box *Box) getVolumn() float64 {
	return box.len * box.width * box.height
}

/*
*
景区门票案例
1) 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于 18，收费 20 元，其它情况门票免
费.
2) 请编写 Visitor 结构体，根据年龄段决定能够购买的门票价格并输出
*/
type Visitor struct {
	Name string
	Age  int
}

func (visitor *Visitor) showPrice() {
	if visitor.Age >= 90 || visitor.Age <= 8 {
		fmt.Println("滚")
		return
	}
	if visitor.Age > 18 {
		fmt.Printf("游客的名字为%v 年龄为%v 收费20元\n", visitor.Name, visitor.Age)
	} else {
		fmt.Printf("游客的名字为%v 年龄为%v 免费\n", visitor.Name, visitor.Age)
	}
}

func main() {
	//测试
	//创建一个Student实例变量
	var stu = Student{
		name:   "toby",
		gender: "male",
		age:    25,
		id:     1000,
		score:  99.5,
	}

	fmt.Println(stu.say())

	var dog = Dog{
		name:   "史努比",
		age:    10,
		weight: 20,
	}
	fmt.Println(dog.dogSay())

	//测试代码
	var box Box
	box.len = 1.1
	box.width = 2.0
	box.height = 3.0

	volumn := box.getVolumn()
	fmt.Printf("体积为=%.2f\n", volumn)

	//测试
	var v Visitor
	for {
		fmt.Println("请输入你的名字")
		fmt.Scanln(&v.Name)
		if v.Name == "n" {
			fmt.Println("退出程序...")
			break
		}
		fmt.Println("请输入你的年龄")
		fmt.Scanln(&v.Age)
		v.showPrice()
	}
}
