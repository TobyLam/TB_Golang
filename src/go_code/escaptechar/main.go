package main

import "fmt" //包中提供格式化、输出、输入函数

// 这是一个main函数,是程序的入口
func main() {
	fmt.Println("welldone")
	fmt.Println("everything will be fine")
	//\t
	fmt.Println("tom\tjack")

	//行注释，快捷键，一次性注释多行 ctrl + /
	//fmt.Println("hello\nworld")
	/*
		块注释
		块注释不可以包含块注释,否则报错
	*/
	/*
		fmt.Println("C:\\Users\\TobyLam\\Desktop\\GolangNote")
		fmt.Println("tom说:\"i love you\"")
	*/

	//\r 回车，从当前行最开始前面开始输出，覆盖掉以前的内容
	fmt.Println("hello,\rworld")
	fmt.Println("天龙八部雪山飞狐\rworld")

	fmt.Println("一行不能有太多字符，保持代码的优雅性，hellowordhellowordhelloword",
		"abcdefghijklmnopqrstuvwxyz",
		"AbcdefghijklmnopqrstuvwxyzAbcdefghi")

}
