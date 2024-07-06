package main

import (
	"encoding/json"
	"fmt"
)

// 定义一个结构体
type Monster struct {
	Name          string `json:"monster_name"` //反射机制
	Age           int    `json:"monster_age"`
	Birthday      string
	Sal           float64
	Skill         string
	monster_alias string //变量名小写，json序列化属于其他包调用，无法生效
}

func testStruct() {
	//演示
	monster := Monster{
		Name:          "牛魔王",
		Age:           500,
		Birthday:      "2011-11-11",
		Sal:           8000.0,
		Skill:         "牛魔拳",
		monster_alias: "公牛",
	}

	//将monster变量序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("序列化错误", err)
	}

	//输出序列化后的结果
	fmt.Printf("monster序列化后=%v\n", string(data))
}

// 将map进行序列化
func testMap() {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "火云洞"

	//将a这个map进行序列化
	//将monster序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	//输出序列化后的结果
	fmt.Printf("a map序列化后=%v\n", string(data))
}

// 对切片进行序列化
func testSlice() {
	var slice []map[string]interface{}
	var m1 map[string]interface{}

	m1 = make(map[string]interface{})

	m1["name"] = "toby"
	m1["age"] = "7"
	m1["address"] = "广州"

	slice = append(slice, m1)

	var m2 map[string]interface{}
	m2 = make(map[string]interface{})

	m2["name"] = "jack"
	m2["age"] = "20"
	m2["address"] = [2]string{"上海", "北京"}

	slice = append(slice, m2)

	//将slice进行序列化
	data, err := json.Marshal(slice)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	//输出序列化后的结果
	fmt.Printf("slice序列化后=%v\n", string(data))
}

// 基本数据类型数据化，对基本数据类型进行序列化，意义不大
func testFloat64() {
	var num1 float64 = 143.42

	//对num1 进行序列化
	data, err := json.Marshal(num1)

	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	//输出序列化后的结果
	fmt.Printf("num1序列化后=%v\n", string(data))
}

func main() {
	//将结构体,map,切片进行序列化
	testStruct()

	testMap()

	testSlice() //切片序列化

	testFloat64() //基本数据类型
}
