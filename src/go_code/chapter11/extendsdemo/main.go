package main
import (
	"fmt"
)

//编写一个学生考试系统

type Student struct{
	Name string
	Age int
	Score int
}

//将Pupil 和 Graduate 共有的方法也绑定到 *Student
func (stu *Student) showInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v\n",stu.Name,stu.Age,stu.Score)
}
//显示成绩
func (stu *Student) SetScore(score int){
	//业务判断
	stu.Score = score
}

//给 *Student 新增一个方法，那么Pupil 和 Graduate 都可以调用该方法
func (stu *Student) GetSum(n1 int,n2 int) int{
	//业务判断
	return n1 + n2
}


//小学生
type Pupil struct{
	Student //嵌入Student匿名结构体
}
//这是Pupil结构体特有的方法，保留
func (p *Pupil) testing(){
	fmt.Println("小学生正在考试中....")
}


//大学生
type Graduate struct{
	Student //嵌入Student匿名结构体
}
//这是Graduate结构体特有的方法，保留
func (g *Graduate) testing(){
	fmt.Println("大学生正在考试中....")
}


//大学生

func main(){
	//测试
	// var pupil = &Pupil{
	// 	Name:"小明",
	// 	Age:12,
	// }
	// pupil.testing()
	// pupil.SetScore(99)
	// pupil.showinfo()

	//当我们对结构体嵌入了匿名结构体，使用方法会发生变化
	pupil := &Pupil{}
	pupil.Student.Name = "小明"
	pupil.Student.Age = 12
	pupil.testing()
	pupil.Student.SetScore(99)
	pupil.Student.showInfo()
	fmt.Println("res=",pupil.Student.GetSum(10,20))

	graduate := &Graduate{}
	graduate.Student.Name = "小张"
	graduate.Student.Age = 22
	graduate.testing()
	graduate.Student.SetScore(99)
	graduate.Student.showInfo()
	fmt.Println("res=",graduate.Student.GetSum(10,200))


}