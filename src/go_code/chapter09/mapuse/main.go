package main

import (
	"fmt"
)

func main() {
	//第一种使用方式
	var a map[string]string
	//在使用map前，需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "赵云" // ok?
	a["no2"] = "孔明" // ok?
	a["no1"] = "陆逊" // ok?
	a["no3"] = "孔明" // ok?
	fmt.Println(a)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"
	fmt.Println(cities)

	//第三种方式
	heroes := map[string]string{
		"hero1": "张良",
		"hero2": "韩信",
		"hero3": "萧何",
	}
	heroes["hero4"] = "白起"
	fmt.Println(heroes)

	//案例
	/**
	课堂练习：演示一个 key-value 的 value 是 map 的案例
	比如：我们要存放 3 个学生信息, 每个学生有 name 和 sex 信息
	思路:map[string]map[string]string
	*/

	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string, 2) //这一行不能少！
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "番禺清河东~"

	studentMap["stu02"] = make(map[string]string, 2)
	studentMap["stu02"]["name"] = "mary"
	studentMap["stu02"]["sex"] = "女"
	studentMap["stu02"]["address"] = "番禺亚运村~"

	fmt.Println(studentMap)
	fmt.Println(studentMap["stu02"])
	fmt.Println(studentMap["stu02"]["address"])

}
