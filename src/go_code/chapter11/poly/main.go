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

//编写一个方法Working 方法，接收一个Usb接口类型的参数
//只要是实现了 Usb接口 ，所谓实现usb接口，就是指实现了Usb接口声明的所有方法
func (c Computer) Working(usb Usb) {

	//通过usb接口变量来调用Start和Stop方法
	usb.Start()
	usb.Stop()
}

func main() {

	//定义一个Usb接口数组，可以存放Phone和Camera的结构体变量
	//多态数组
	var usbArr [3]Usb
	usbArr[0] = Phone{"vivo"}
	usbArr[1] = Phone{"小米"}
	usbArr[2] = Camera{"尼康"}

	fmt.Println(usbArr)
}
