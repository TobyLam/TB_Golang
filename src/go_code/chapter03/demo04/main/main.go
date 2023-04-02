package main

import (
  "fmt"
  //为了使用utils.go文件的变量或者函数，我们需要先引进该model包
  "go_code/chapter03/demo04/model"
)

// 变量使用的注意事项
func main() {

	//该区域的数据值可以在同一类型范围内不断变化
	var i int
	i = 2
	i = 3
	fmt.Println("i=", i)
	//i = 1.2 //报错，原因是不能改变数据类型

	//变量在同一个作用域(在一个函数或者在代码块)内不能重名
	//var i int = 99 //报错，重复定义
	//i := 99

	//我们使用utils.go 的 heroName 包名.标识符
	fmt.Println(model.HeroName)

}
