package main

import "fmt"

type Binterface interface {
	test01()
}

type CInterface interface {
	test02()
	test01() //低版本不支持两个同名方法,AInterface接口中只能有一个同名方法
}

type AInterface interface {
	Binterface
	CInterface
	test03()
}

// 如果需要实现AInterface接口，需要实现Binterface和CInterface接口
type Stu struct {
}

func (stu Stu) test01() {

}
func (stu Stu) test02() {

}
func (stu Stu) test03() {

}

type T interface { //空接口
}

type Usb interface {
	Say()
}

type Stt struct {
}

// func (this *Stt) Say() {
// 	fmt.Println("Say()~")
// }
func (stt *Stt) Say() {
	fmt.Println("Say()~")
}

func main() {
	var stu Stu
	var a AInterface = stu
	a.test01()

	var t T = stu //ok
	fmt.Println(t)
	var t2 interface{} = stu //ok
	var num1 float64 = 10.1
	t2 = num1
	t = num1
	fmt.Println(t2)

	var stt Stt = Stt{}
	// 错误!会报Stt类型没有实现接口Usb
	// var sttu Usb = stt
	var sttu Usb = &stt //ok
	sttu.Say()
	fmt.Println("here", sttu)

}
