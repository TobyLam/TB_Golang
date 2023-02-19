package main

import (
	"fmt"
)

// 演示golang中string 类型的使用
func main() {
	//string的基本使用
	var address string = "北京长城 110 hello world"
	fmt.Println(address)

	//字符串一旦赋值了， 字符串就不能修改了： 在 Go 中字符串是不可变的
	//var str = "hello"
	//str[0] = 'a' //这里就不能去修改str的内容，即go中的字符串是不可变的

	//字符串的两种表示形式
	//(1) 双引号, 会识别转义字符
	//(2) 反引号， 以字符串的原生形式输出， 包括换行和特殊字符， 可以实现防止攻击、
	//输出源代码等效果 【案例演示】
	str2 := "abc\nabc"
	fmt.Println(str2)

	//使用的反引号 ``
	str3 := `
	package main

	import (
		"fmt"
	"unsafe"
	)

	// 演示golang中bool类型的使用
	func main() {
		var b = false
		fmt.Println("b=", b)
		//注意事项
		//1.bool类型占用存储空间是1个字节
		fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
		//2.bool类型的取值只能是true 或 false
		//b = 1 错误
	}
	`
	fmt.Println(str3)

	//字符串拼接方式
	var str = "hello " + "world"
	str += " haha!"
	fmt.Println(str)

	//当一个拼接操作很长时，可以分行写,但是注意，需要将+号保留在上一行
	str4 := "hello " + "world" + "hello " +
		"world" + "hello " + "world" + "hello " + "world" +
		"hello " + "world" + "hello " + "world" + "hello " +
		"world" + "hello " + "world" + "hello " + "world" + "hello " +
		"world" + "hello " + "world" + "hello " + "world"
	fmt.Println(str4)

	//数据类型默认值
	var a int          // 0
	var b float32      // 0
	var c float64      // 0
	var isMarried bool //false
	var name string    //""
	//这里的%v表示按照变量的值输出
	fmt.Printf("a=%d,b=%f,c=%f isMarried=%v name=%v", a, b, c, isMarried, name)
}
