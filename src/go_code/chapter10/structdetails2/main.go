package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}

type B struct {
	Num int
}

type B1 struct {
	Num float64
}

type B2 struct {
	Num  int
	Name string
}

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	a = A(b) // ? 可以转换，但是有要求，就是结构体的字段要完全一样（包括：名字、个数和类型）
	//a = B1(b)  //报错
	//a = B2(b)  //报错
	fmt.Println(a, b)

	//1、创建一个Monster变量
	monster := Monster{"牛魔王", 500, "芭蕉扇~"}

	//2、将monster变量序列化为json格式字串
	// json.Marshal 函数中使用反射，这个讲解反射时，详细介绍
	jsonMonster, err := json.Marshal(monster)
	if err != nil {
		fmt.Println("json 处理错误 ", err)
	}
	fmt.Println("jsonStr", string(jsonMonster))
}
