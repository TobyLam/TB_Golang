package main

import (
	"fmt"
)

type MethodUtils struct {
	//字段
}

// 给MethodUtils编写方法
func (mu MethodUtils) Print() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 8; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 2) 编写一个方法，提供 m 和 n 两个参数，方法中打印一个 m*n 的矩形
func (mu MethodUtils) Print2(m int, n int) {
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// 3) 编写一个方法算该矩形的面积(可以接收长 len，和宽 width)， 将其作为方法返回值。在 main
// 方法中调用该方法，接收返回的面积值并打印。
func (mu MethodUtils) area(len float64, width float64) float64 {
	return len * width
}

// 4) 编写方法：判断一个数是奇数还是偶数
func (mu *MethodUtils) JudgeNum(num int) {
	if num%2 == 0 {
		fmt.Println(num, "是偶数。。。")
	} else {
		fmt.Println(num, "是奇数。。。")
	}
}

// 5) 根据行、列、字符打印 对应行数和列数的字符，比如：行：3，列：2，字符*,则打印相应的效
// 果
func (mu *MethodUtils) Print3(n int, m int, key string) {
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			fmt.Print(key)
		}
		fmt.Println()
	}
}

// 6) 定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
// 实现形式 1：分四个方法完成: ,分别计算 +-*/
// 实现形式 2：用一个方法搞定,需要接收两个数，还有一个运算符
type Calcuator struct {
	Num1 float64
	Num2 float64
}

func (calcuator *Calcuator) getSum() float64 {
	return calcuator.Num1 + calcuator.Num2
}

func (calcuator *Calcuator) getSub() float64 {
	return calcuator.Num1 - calcuator.Num2
}

func (calcuator *Calcuator) getRes(operator byte) float64 {
	res := 0.0
	switch operator {
	case '+':
		res = calcuator.Num1 + calcuator.Num2
	case '-':
		res = calcuator.Num1 - calcuator.Num2
	case '*':
		res = calcuator.Num1 * calcuator.Num2
	case '/':
		res = calcuator.Num1 / calcuator.Num2
	default:
		fmt.Println("运算符输入有误...")
	}
	return res
}

// (7)在MethodUtils结构体编个方法，从简谱接收整数（1-9），打印对应乘法表；
func (mu *MethodUtils) Print4() {
	num := 0
	fmt.Println("请输入一个整数（1-9）")
	fmt.Scanln(&num)
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%v × %v = %v ", i, j, i*j)
		}
		fmt.Println()
	}
}

// (8)编写方法，使给定的一个二维数组（3×3）转置：
func (mu *MethodUtils) Print5() {
	var arr [3][3]int
	var arr2 [3][3]int
	val := 0
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			val++
			arr[i][j] = val
		}
	}

	fmt.Println("转置前：arr=", arr)

	//转置
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			arr2[i][j] = arr[j][i]
		}
	}
	fmt.Println("转置后arr=", arr2)
}
func main() {
	//1) 编写结构体(MethodUtils)，编程一个方法，方法不需要参数，在方法中打印一个 10*8 的矩形，
	//在 main 方法中调用该方法。
	var mu MethodUtils
	mu.Print()
	fmt.Println()
	mu.Print2(5, 20)

	fmt.Println()
	areaRes := mu.area(2.5, 8.7)
	fmt.Println("面积为", areaRes)

	fmt.Println()

	(&mu).JudgeNum(10)
	mu.JudgeNum(11)

	mu.Print3(7, 20, "@")

	//测试一下
	var calcuator Calcuator
	calcuator.Num1 = 1.2
	calcuator.Num2 = 2.2

	fmt.Println("sum=", fmt.Sprintf("%.2f", calcuator.getSum()))
	fmt.Println("sub=", fmt.Sprintf("%.2f", calcuator.getSub()))

	res := calcuator.getRes('+')
	fmt.Println("res=", res)

	//打印乘法表
	mu.Print4()

	//二维数组转置
	mu.Print5()
}
