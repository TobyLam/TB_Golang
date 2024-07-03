package main

import (
	"fmt"
	"os"
)

func main() {

	//打开一个文件
	//概念说明：file的叫法
	//1.file指针
	//2.file对象
	//3.file 文件句柄
	file, err := os.Open("f:/go_workspace/TB_Golang/testdd.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//输出文件，看看文件是什么 ==> <指针*File>
	fmt.Printf("file=%v", file)

	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file err=", err)
	}
}
