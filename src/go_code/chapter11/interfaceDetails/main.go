package main

import "fmt"

type AInterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println("Stu Say()")
}

type interger int

func (i interger) Say() {
	fmt.Println("interger Say() i=", i)
}

type BInterface interface {
	Hello()
}

type Monster struct {
}

func (m Monster) Hello() {
	fmt.Println("Monster Hello()")
}
func (m Monster) Say() {
	fmt.Println("Monster Say()~")
}

func main() {

	var stu Stu //结构体变量，实现了 Say() ,实现了 Ainterface 接口
	var a AInterface = stu
	a.Say()

	var i interger = 10
	var b AInterface = i
	b.Say()

	//Monster实现了AInterface 和 BInterface 接口
	var monster Monster
	var a2 AInterface = monster
	var b2 BInterface = monster
	a2.Say()
	b2.Hello()
}
