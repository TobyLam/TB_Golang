package main

import "fmt"

//声明/定义一个接口
type Usb interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
}

//声明/定义一个接口
type Usb2 interface {
	//声明了两个没有实现的方法
	Start()
	Stop()
	Test()
}

type Phone struct {
	name string
}

//让Phone实现Usb接口的方法
func (p Phone) Start() {
	fmt.Println("手机开始工作。。。")
}
func (p Phone) Stop() {
	fmt.Println("手机停止工作。。。")
}

func (p Phone) Call() {
	fmt.Println("手机在打电话。。。")
}

//让Camera实现Usb接口的方法
type Camera struct {
	name string
}

func (c Camera) Start() {
	fmt.Println("相机开始工作~~~。。。")
}
func (c Camera) Stop() {
	fmt.Println("相机停止工作。。。")
}

//计算机
type Computer struct {
}

func (computer Computer) Working(usb Usb) {
	usb.Start()
	//如果usb是指向Phone结构体变量，则还需要调用Phone特有的方法call
	//类型断言
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {

	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	//遍历usbArr
	//Phone 还有一个特有的方法 call()，请遍历 Usb 数组，如果是 Phone 变量，
	//除了调用 Usb 接口声明的方法外，还需要调用 Phone 特有方法 call. =》类型断言
	var computer Computer
	for _, v := range usbArr {
		computer.Working(v)
		fmt.Println()
	}
	fmt.Println(usbArr)
}
