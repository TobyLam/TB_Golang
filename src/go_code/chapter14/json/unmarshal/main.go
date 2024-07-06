package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Age      int
	Birthday string
}

type Dog struct {
	Age      int
	Birthday string
}

// 将map进行序列化
func testMap() string {
	//定义一个map
	var a map[string]interface{}
	//使用map,需要make
	a = make(map[string]interface{})
	a["Age"] = "红孩儿~~~~"
	a["Birthday"] = "2024-07-06"

	//将a这个map进行序列化
	//将monster序列化
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}

	//输出序列化后的结果
	return string(data)
}

// 将json字符串，反序列化成struct
func unmarshalStruct() {
	str := "{\"Age\":500,\"Birthday\":\"2024-07-06\"}"

	//定义一个Monster实例
	var monster Dog
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("unmarshal err=%v\n", err)
	}
	fmt.Printf("反序列化后 monster=%v\n monster.Age=%v", monster, monster.Age)
}

// 将json字符串，反序列化成map
func unmarshalMap() {
	str := testMap()

	//定义一个map
	var a map[string]interface{}

	//反序列化
	//注意：反序列化map，不需要make,因为make操作被封装到 Unmarshal函数
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}

	fmt.Printf("反序列化后 a=%v\n", a)
}

// 将json字符串，反序列化成切片
func unmarshalSlice() {
	str := "[{\"address\":\"广州\",\"age\":\"7\",\"name\":\"toby\"}," +
		"{\"address\":[\"上海\",\"北京\"],\"age\":\"20\",\"name\":\"jack\"}]"

	//定义一个slice
	var a []map[string]interface{}

	//反序列化
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("unmarshal err=%v", err)
	}

	fmt.Printf("反序列化后 a=%v\n", a)

}
func main() {
	unmarshalStruct()
	unmarshalMap()
	unmarshalSlice()
}
