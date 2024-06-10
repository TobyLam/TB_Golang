package main

import "fmt"

func main() {
	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "秦皇岛"
	cities["no2"] = "铁岭"
	cities["no3"] = "攀枝花"
	fmt.Println(cities)

	//因为no3这个key已经存在，因此下面这句话就相当于修改
	cities["no3"] = "宜春"

	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作,也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//演示map的查找
	val, ok := cities["no2"]
	if ok {
		fmt.Printf("有no2 key 值为%v\n", val)
		fmt.Println(ok)
	} else {
		fmt.Printf("没有no1 key\n")
	}

	//如果希望一次性删除所有的key
	//1.遍历所有的key，如何逐一删除[遍历]
	//2.直接make一个新的空间
	cities = make(map[string]string)
	fmt.Println(cities)
}
