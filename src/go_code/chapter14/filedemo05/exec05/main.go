package main

import (
	"fmt"
	"os"
)

func main() {
	//将f:/go_workspace/TB_Golang/hello.txt 文件内容导入到 f:/go_workspace/TB_Golang/test.txt

	//1.将f:/go_workspace/TB_Golang/hello.txt 文件内容读取到内存
	//2.将内存中的内容写入到 f:/go_workspace/TB_Golang/test.txt

	file1Path := "f:/go_workspace/TB_Golang/hello.txt"
	file2Path := "f:/go_workspace/TB_Golang/test.txt"

	data, err := os.ReadFile(file1Path)
	if err != nil {
		//读取文件错误
		fmt.Println("read file err=", err)
		return
	}

	err = os.WriteFile(file2Path, data, 0666)

	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}

}
