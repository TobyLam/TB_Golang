package main
import (
	"fmt"
	"go_code/chapter11/encapsulate/model"
)
/**
请大家看一个程序(person.go),不能随便查看人的年龄,工资等隐私，并对输入的年龄进行合理的验
证。设计: model 包(person.go) main 包(main.go 调用 Person 结构体)
*/
func main(){
	
	p := model.NewPerson("toby")

	p.SetAge(20)
	p.SetSal(10000)

	fmt.Println(p)
	fmt.Println(p.Name, " age =",p.GetAge(), " sal =",p.GetSal())
}