package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//统计字符串的长度，按字节 len(str)
	//golang的编码统一为utf-8 (ascii的字符（字母和数字）占一个字节，汉字占用3个字节)
	str := "hello广"
	fmt.Println("str len=", len(str)) //8

	str2 := "hello广州"
	//字符串遍历，同时处理有中文的问题 r := []rune(str)
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}
	//字符串转整数: n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println("转换错误", err)
	} else {
		fmt.Println("转成的结果是", n)
	}

	//4) 整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v,str=%T\n", str, str)

	//5) 字符串 转 []byte: var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//6) []byte 转 字符串: str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Printf("str=%v\n", str)

	//10 进制转 2, 8, 16 进制: str = strconv.FormatInt(123, 2),返回对应的字符串 // 2-> 8 , 16
	str = strconv.FormatInt(123, 2)
	fmt.Printf("123对应的二进制是=%v\n", str)
	str = strconv.FormatInt(123, 16)
	fmt.Printf("123对应的十六进制是=%v\n", str)

	//8) 查找子串是否在指定的字符串中: strings.Contains("seafood", "foo") //true
	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	//9) 统计一个字符串有几个指定的子串 ： strings.Count("ceheese", "se") //4
	num := strings.Count("ceheese", "se")
	fmt.Printf("num=%v\n", num)

	// 10) 不区分大小写的字符串比较(==是区分字母大小写的): fmt.Println(strings.EqualFold("abc", "Abc")) // true
	b = strings.EqualFold("abc", "Abc")
	fmt.Printf("b=%v\n", b) //true

	fmt.Println("结果", "abc" == "Abc") // false //区分字母大小写

	//11) 返回子串在字符串第一次出现的 index 值，如果没有返回-1 :
	// strings.Index("NLT_abc", "abc") // 4
	index := strings.Index("NLT_abcabcabc", "abc")
	fmt.Printf("index=%v\n", index) //4

	//12)返回子串在字符串最后一次出现的 index，
	// 如没有返回-1 : strings.LastIndex("go golang", "go")
	index = strings.LastIndex("go golang", "go")
	fmt.Printf("index=%v\n", index) // 3

	//13) 将指定的子串替换成 另外一个子串: strings.Replace("go go hello", "go", "go 语言", n) n 可以指
	//定你希望替换几个，如果 n=-1 表示全部替换
	str2 = "go go hello"
	str = strings.Replace(str2, "go", "广州", -1)
	fmt.Printf("str=%v str2=%v\n", str, str2)

	//14) 按 照 指 定 的 某 个 字 符 ， 为 分 割 标 识 ， 将 一 个 字 符 串 拆 分 成 字 符 串 数 组 ：
	//strings.Split("hello,wrold,ok", ",")
	strArr := strings.Split("hello,world,ok", ",")
	for i := 0; i < len(strArr); i++ {
		fmt.Printf("str[%v]=%v\n", i, strArr[i])
	}
	fmt.Printf("strArr=%v\n", strArr)

	//15) 将字符串的字母进行大小写的转换: strings.ToLower("Go") // go strings.ToUpper("Go") // GO
	str = "goLang Hello"
	str = strings.ToLower(str)
	str = strings.ToUpper(str)
	fmt.Printf("str=%v\n", str)

	//16) 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntrn ")
	str = strings.TrimSpace(" tn a lone gopher ntrn ")
	fmt.Printf("str=%q\n", str)

	//17) 将字符串左右两边指定的字符去掉 ： strings.Trim("! hello! ", " !") // ["hello"] //将左右两边 !
	//和 " "去掉
	str = strings.Trim("! he!llo! ", " !")
	fmt.Printf("str=%v\n", str)

	//18) 将字符串左边指定的字符去掉 ： strings.TrimLeft("! hello! ", " !") // ["hello"] //将左边 ! 和 "
	//"去掉
	//19) 将字符串右边指定的字符去掉 ：strings.TrimRight("! hello! ", " !") // ["hello"] //将右边 ! 和 "
	//"去掉

	//20) 判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1", "ftp") // true
	b = strings.HasPrefix("ftp://192.168.10.1", "ftp")
	fmt.Printf("b=%v\n", b)

	//21) 判断字符串是否以指定的字符串结束: strings.HasSuffix("NLT_abc.jpg", "abc") //false
}
