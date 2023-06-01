package main

import (
	"fmt"
)

// 写一个非常简单的函数
func test(char byte) byte {
	return char + 1
}

func main() {
	//案例：
	//请编写一个程序，该程序可以接收一个字符，比如: a,b,c,d,e,f,g a表示星期一，b表示星期二 … 根
	//据用户的输入显示相依的信息.要求使用 switch 语句完成

	//思路分析
	// 1.定义一个变量接收字符
	// 2.使用switch完成

	//var key byte
	//fmt.Println("请输入一个字符 a,b,c,d,e,f,g")
	//fmt.Scanf("%c", &key)
	//
	//switch test(key) + 1 { //语法现象
	//case 'a':
	//	fmt.Println("周一,猴子穿新衣")
	//case 'b':
	//	fmt.Println("周二，猴子当小二")
	//case 'c':
	//	fmt.Println("周三，猴子爬雪山")
	////...
	//default:
	//	fmt.Println("输入有误，猴子在忙...")
	//}

	//var n1 int32 = 20
	//var n2 int32 = 20
	////var n3 int32 = 5
	//switch n1 {
	//case n2, 10, 5: // 错误，原因是 n2的数据类型和n1不一致
	//	fmt.Println("ok1")
	////case 5: //错误，因为前面我们有常量5，因此重复，就会报错
	////	fmt.Println("ok2")
	//default:
	//	fmt.Println("没有匹配到")
	//}

	//switch 后也可以不带表达式，类似 if --else 分支来使用。
	var age int = 20

	switch {
	case age == 10:
		fmt.Println("age == 10")
	case age == 20:
		fmt.Println("age == 20")
	default:
		fmt.Println("没有匹配到")
	}

	//case 中也可以对 范围进行判断
	var score int = 90
	switch {
	case score > 90:
		fmt.Println("成绩优秀..")
	case score >= 70 && score <= 90:
		fmt.Println("成绩优良...")
	case score >= 60 && score < 70:
		fmt.Println("成绩及格...")
	default:
		fmt.Println("不及格")
	}

	//switch 后也可以直接声明/定义一个变量，分号结束，不推荐
	switch grade := 90; { //在golang中，可以这样写
	case grade > 90:
		fmt.Println("成绩优秀~..")
	case grade >= 70 && grade <= 90:
		fmt.Println("成绩优良~...")
	case grade >= 60 && grade < 70:
		fmt.Println("成绩及格~...")
	default:
		fmt.Println("不及格")
	}

	//switch 的穿透 fallthrough
	var num int = 10
	switch num {
	case 10:
		fmt.Println("ok1")
		fallthrough //默认只能穿透一层
	case 20:
		fmt.Println("ok2")
	case 30:
		fmt.Println("ok3")
	default:
		fmt.Println("没有匹配到..")
	}

	//Type Switch：switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际指向的
	//变量类型
	var x interface{}
	var y = 10.0
	x = y
	switch i := x.(type) {
	case nil:
		fmt.Printf("x 的类型~ ： %T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println(" x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型")
	default:
		fmt.Printf("未知型")
	}
}
