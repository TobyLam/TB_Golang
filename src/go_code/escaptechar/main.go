package main

import "fmt" //包中提供格式化、输出、输入函数

func main() {
	//\t
	fmt.Println("tom\tjack")
	//\n
	fmt.Println("hello\nworld")
	//\\
	fmt.Println("C:\\Users\\TobyLam\\Desktop\\GolangNote")
	//\"
	fmt.Println("tom说:\"i love you\"")
	//\r 回车，从当前行最开始前面开始输出，覆盖掉以前的内容
	fmt.Println("hello,\rworld")
	fmt.Println("天龙八部雪山飞狐\rworld")

}
