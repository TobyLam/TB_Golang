package main

import (
	"fmt"
)

func main() {
	//map的声明和注意事项

	var a map[string]string
	//在使用map前，需要先make,make的作用就是给map分配数据空间
	a = make(map[string]string, 10)
	a["no1"] = "赵云" // ok?
	a["no2"] = "孔明" // ok?
	a["no1"] = "陆逊" // ok?
	a["no3"] = "孔明" // ok?

	fmt.Println(a)
}
