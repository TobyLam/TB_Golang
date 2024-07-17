package main

import (
	"fmt"
	"reflect"
)

// 专门演示反射
func reflectTest01(b interface{}) {

	//通过反射获取到传入变量的type、kind、值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)
	//n1 := 10
	//n2 := 2 + rVal //error
	n2 := 2 + rVal.Int()
	//n3 := rVal.Float() //error ,要求数据类型匹配
	fmt.Println("n2=", n2)
	//fmt.Println("n3=", n3)

	fmt.Printf("rVal=%v rVal type=%T\n", rVal, rVal)

	//将rVal 转成 interface{}
	iV := rVal.Interface()

	//将 interFace{} 通过断言转成需要的类型
	num2 := iV.(int)
	fmt.Println("num2=", num2)
}

// 专门演示反射[对结构体的反射]
func reflectTest02(b interface{}) {
	//通过反射获取到传入变量的type、kind、值
	//1.先获取到reflect.Type
	rTyp := reflect.TypeOf(b)
	fmt.Println("rType=", rTyp)

	//2.获取到 reflect.Value
	rVal := reflect.ValueOf(b)

	//3.获取 变量对应的Kind
	//(1) rVal.Kind() ==>
	kind1 := rVal.Kind()
	//(2) rTyp.Kind() ==>
	kind2 := rTyp.Kind()
	fmt.Printf("kind =%v kind=%v\n", kind1, kind2)

	//将rVal 转成 interface{}
	iV := rVal.Interface()
	fmt.Printf("iV = %v iV=%T\n", iV, iV)

	//将 interFace{} 通过断言转成需要的类型
	//带检测的类型断言
	//可以使用switch的断言形式做的更加灵活
	stu, ok := iV.(Student)
	if ok {
		fmt.Printf("stu.Name=%v\n", stu.Name)
	}
}

type Student struct {
	Name string
	Age  int
}

type Monster struct {
	Name string
	Age  int
}

func main() {
	//案例：
	//对(基本数据类型、interface{}、reflect.Value)进行反射的基本操作

	//1.先定义一个int
	var num int = 100
	reflectTest01(num)

	//2.定义一个Student的实例
	stu := Student{
		Name: "tom",
		Age:  20,
	}
	reflectTest02(stu)
}
