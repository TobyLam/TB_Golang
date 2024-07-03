package main

import (
	"fmt"
	_ "io/ioutil"
	"os"
)

func main() {
	//使用ioutil.Readfile一次性将文件读取到位【文件不能太大】
	file := "f:/go_workspace/TB_Golang/test.txt"
	//ioutil.ReadFile(file)
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("read file err=%v\n", err)
	}
	//把读取到的内容显示到终端
	fmt.Printf("%v\n", string(content)) // []byte

	//因为,没有显式的Open文件,所以,不需要显式的Close文件
	//因为，文件的Open和Close被封装到ReadFile函数中
}
