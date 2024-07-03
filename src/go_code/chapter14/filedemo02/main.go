package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	//打开文件
	file, err := os.Open("f:/go_workspace/TB_Golang/test.txt")
	if err != nil {
		fmt.Println("open file err=", err)
	}

	//当函数退出时，要及时关闭file
	defer file.Close() //要及时关闭file句柄，否则会有内存泄漏

	// 创建一个 *Reader ,是带缓冲的
	/**
	const (
	defaultBufSize = 4096 //默认的缓冲区大小
	)
	*/
	reader := bufio.NewReader(file)

	//循环读取文件
	for {
		str, err := reader.ReadString('\n') //读到一个换行就结束
		if err == io.EOF {                  //io.EOF 表示文件的末尾
			break
		}
		//输出内容
		fmt.Print(str)
	}

	fmt.Println("文件读取结束...")
}
