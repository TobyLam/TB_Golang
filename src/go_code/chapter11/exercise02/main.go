package main

import (
	"fmt"
	"math/rand"
	"sort"
)

// 1.声明Hero结构体
type Hero struct {
	Name string
	Age  int
}

type Student struct {
	Name  string
	Age   int
	Score float64
}

// 2.声明一个Hero结构体切片类型
type HeroSlice []Hero

type StudentSlice []Student

// 3.实现HeroSlice的Len()，Swap()，Less()方法
// 实现Interface接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

func (hs HeroSlice) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

// Less()方法决定使用什么标准进行排序
// 1.按照年龄从小到大排序
func (hs HeroSlice) Less(i, j int) bool {
	// return hs[i].Age < hs[j].Age
	// 2.按照姓名的字母顺序排序
	return hs[i].Name < hs[j].Name
}

// 4.实现StudentSlice的Len()，Swap()，Less()方法
func (ss StudentSlice) Len() int {
	return len(ss)
}

func (ss StudentSlice) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}

// Less()方法决定使用什么标准进行排序
// 1.按照年龄从小到大排序
func (ss StudentSlice) Less(i, j int) bool {
	// return ss[i].Age < ss[j].Age
	// 2.按照成绩从大到小排序
	return ss[i].Score > ss[j].Score
}

func main() {
	//先定义一个数组/切片
	var intSlice = []int{0, -1, 10, 7, 90}
	//要求对 intSlice 切片进行排序
	//1. 冒泡排序...
	//2.也可以使用系统提供的方法
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	//对结构体切片进行排序
	//1.冒泡排序...
	//2.使用系统提供的方法

	//测试对结构体切片进行排序
	var heroes HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("英雄%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		//将hero添加到切片中append
		heroes = append(heroes, hero)
	}

	//打印排序前的切片
	for _, v := range heroes {
		fmt.Println(v)
	}
	fmt.Println("==================排序后==================")
	//调用sort.Sort()方法进行排序
	sort.Sort(heroes)
	//打印排序后的切片
	for _, v := range heroes {
		fmt.Println(v)
	}

	//测试对结构体切片进行排序
	var students StudentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			Name:  fmt.Sprintf("学生%d", rand.Intn(100)),
			Age:   rand.Intn(100),
			Score: rand.Float64() * 100,
		}
		//将hero添加到切片中append
		students = append(students, student)
	}

	//打印排序前的切片
	for _, v := range students {
		fmt.Println(v)
	}
	fmt.Println("==================排序后==================")
	//调用sort.Sort()方法进行排序
	sort.Sort(students)
	//打印排序后的切片
	for _, v := range students {
		fmt.Println(v)
	}

}
