package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	A
	B
	//Name string
}

type D struct {
	a    A //有名结构体 [组合关系]
	Name string
}

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

// 多重继承
type TV struct {
	Goods
	Brand
}

type TV2 struct {
	*Goods
	*Brand
}

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int //匿名字段，基本数据类型
	n   int
}

func main() {
	var c C
	//如果c没有Name字段，而 A 和B 有Name,这时就必须通过指定匿名结构体名字来区分
	//所有 c.Name 就会报编译错误，这个规则对方法也适用
	c.A.Name = "tom" //error
	fmt.Println(c)

	//如果D中是一个有名结构体，则访问有名结构体字段时，必须带上结构体名字
	//比如d.a.Name
	var d D
	d.a.Name = "jack"
	d.Name = "mary"

	//嵌套匿名结构体后，也可以在创建结构体变量(实例)时，直接指定匿名结构体字段的值
	tv := TV{
		Goods{"电视机001", 5000},
		Brand{"长虹", "四川"},
	}

	//演示访问Goods的Name
	fmt.Println(tv.Goods.Name)
	fmt.Println(tv.Price)

	tv2 := TV{
		Goods{
			Name:  "电视机002",
			Price: 6000,
		},
		Brand{
			Name:    "海信",
			Address: "山东",
		},
	}
	tv3 := TV2{
		&Goods{"电视机003", 7000},
		&Brand{"小米", "北京"},
	}
	tv4 := TV2{
		&Goods{
			Name:  "电视机004",
			Price: 8000,
		},
		&Brand{
			Name:    "格力",
			Address: "广东",
		},
	}
	fmt.Println("tv", tv)
	fmt.Println("tv2", tv2)
	fmt.Println("tv3", *tv3.Goods, *tv3.Brand)
	fmt.Println("tv4", *tv4.Goods, *tv4.Brand)

	//演示一下匿名字段是基本数据类型的使用
	var e E
	e.Name = "牛魔王"
	e.Age = 500
	e.int = 20
	e.n = 30
	fmt.Println("e=", e)

}
