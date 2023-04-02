package main

import (
	"fmt"
)

//演示golang中标识符的使用
func main(){
	//Golang中严格区分大小写
	//golang 中认为 num 和 Num是两个不同的变量
	var num int = 10
	var Num int = 20

	fmt.Printf("num=%v Num=%v",num,Num)

	//标识符不能保护空格
	//var ab c int = 30

	// _ 是空标识符， 用于占位
	//var _ int = 40 //error
	//fmt.Println(_)

	//var if int = 10 //保留字不能用作标识符

	//语法可以通过，要求不这样使用
	var int int = 20
	var float32 float32 = 12.12345
	fmt.Printf("int Type=%T float32 Type=%T",int,float32)
}