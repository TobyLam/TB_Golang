package main

import (
	"fmt"
	"go_code/chapter10/factory/model"
)

func main() {
	//创建要给Student实例
	//var stu = model.student{
	//	Name:  "tom",
	//	Score: 78.9,
	//}

	//当student结构体是首字母小写，可以通过工厂模式来解决
	var stu = model.NewStudent("tom~", 98.8)
	fmt.Println(*stu) // &{}
	fmt.Println("name=", stu.Name, "score=", stu.GetScore())
}
