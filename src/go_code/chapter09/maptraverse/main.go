package main

import (
	"fmt"
)

func main() {
	//使用for-range遍历map
	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "泰安"
	cities["no2"] = "绵阳"
	cities["no3"] = "岳阳"

	for k, v := range cities {
		fmt.Printf("k=%v v=%v\n", k, v)
	}

	fmt.Println("cities 有", len(cities), "对key-value")

	//使用for-range遍历一个结构较复杂的map
	studentMap := make(map[string]map[string]string)

	studentMap["stu01"] = make(map[string]string, 3)
	studentMap["stu01"]["name"] = "姜维"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "甘肃天水~"

	studentMap["stu02"] = make(map[string]string, 3)
	studentMap["stu02"]["name"] = "赵云"
	studentMap["stu02"]["sex"] = "男"
	studentMap["stu02"]["address"] = "河北常山~"

	for k1, v1 := range studentMap {
		fmt.Println("k1=", k1)
		for k2, v2 := range v1 {
			fmt.Printf("\t k2=%v, v2=%v\n", k2, v2)
		}
		fmt.Println()
	}
}
