package main

import (
	"fmt"
	"reflect"
)

// 1）编写一个结构体，有两个字段 Num1,和Num2
type Cal struct {
	Num1 int
	Num2 int
}

// 2)方法GetSub(name string)
func (c Cal) GetSub(name string) {
	fmt.Printf("%v 完成了减法运行，%v-%v=%v\n", name, c.Num1, c.Num2, c.Num1-c.Num2)
}

// 3)使用反射遍历Cal结构体所有的字段信息

// 4)使用反射机制完成对GetSub的调用，输出形式为
// "tom 完成了减法运行，8-3=5

func main() {

	//创建一个Cal实例
	var a Cal

	//取得 reflect.Value
	val := reflect.ValueOf(a)

	//获取类别，如果不是结构体就退出
	kd := val.Kind()
	if kd != reflect.Struct {
		fmt.Println("struct expect")
		return
	}
	//获取字段个数
	num := val.NumField()
	fmt.Printf("struct has %d Fields\n", num)
	//3)使用反射遍历Cal结构体所有的字段信息
	for i := 0; i < num; i++ {
		fmt.Printf("Field %d 值等于=%v\n", i, val.Field(i))
	}
	rVal := reflect.ValueOf(&a)
	//调用
	rVal.Elem().FieldByName("Num1").SetInt(8)
	rVal.Elem().FieldByName("Num2").SetInt(3)

	var params []reflect.Value
	params = append(params, reflect.ValueOf("tom"))
	rVal.Elem().Method(0).Call(params)

}
