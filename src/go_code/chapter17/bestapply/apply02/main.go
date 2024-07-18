package main

import (
	"fmt"
	"reflect"
)

// 定义Monster结构体
type Monster struct {
	Name    string  `json:"name"`
	Age     int     `json:"monster_age"`
	Score   float32 `json:"成绩"`
	Gender  string
	Address string
}

// 方法，显示s的值
func (s Monster) Print() {
	fmt.Println("---start---")
	fmt.Println(s)
	fmt.Println("---end---")
}

// 方法，返回两个数的和
func (s Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// 方法，给s赋值
func (s Monster) Set(name string, age int, score float32, gender string) {
	s.Name = name
	s.Age = age
	s.Score = score
	s.Gender = gender
}

func TestStruct(a interface{}) {
	//获取 reflect.Type 类型
	typ := reflect.TypeOf(a)
	//获取 reflect.Value 类型
	val := reflect.ValueOf(a)
	//获取 a 对应的类别
	kd := val.Kind()
	//如果传入的不是地址就退出
	if kd != reflect.Ptr && val.Elem().Kind() == reflect.Struct {
		fmt.Println("expect struct")
		return
	}
	//获取到有几个字段
	num := val.Elem().NumField()
	//修改第一个字段的值
	val.Elem().Field(0).SetString("江小白")
	//遍历字段的值
	for i := 0; i < num; i++ {
		fmt.Printf("%d %v\n", i, val.Elem().Field(i).Kind())
	}
	fmt.Printf("struct has %d fields\n", num)

	tag := typ.Elem().Field(0).Tag.Get("json")
	fmt.Printf("tag=%s\n", tag)

	numOfMethod := val.Elem().NumMethod()
	fmt.Printf("struct has %d methods\n", numOfMethod)

	//val.Elem().Method(0).Call(nil)
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(20))
	res := val.Elem().Method(0).Call(params)

	fmt.Println("res=", res[0].Int())

}

//2)使用反射的方式来获取结构体的tag标签，遍历字段的值，修改字段值，调用结构体方法（要求：
//通过传递地址的方式完成，在前面的案例上修改即可）

func main() {
	//创建一个实例
	a := Monster{
		Name:    "国窖2024",
		Age:     200,
		Score:   99,
		Gender:  "女",
		Address: "北京",
	}

	TestStruct(&a)

	fmt.Println(a)
}
