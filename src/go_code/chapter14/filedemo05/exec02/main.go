package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//打开一个存在的文件中，将原来的内容覆盖成新的内容 10 句 "你好，golang!"

	//创建一个新文件，写入内容 5句"hello,Gardon"
	//1.打开文件已经存在文件 f:/go_workspace/TB_Golang/hello.txt
	filePath := "f:/go_workspace/TB_Golang/hello.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	//准备写入10句"你好,golang"
	str := "你好,golang\r\n" // \n表示换行
	//写入时，使用带缓存的 *Writer
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	//因为writer是带缓存的，因此在调用WriterString方法时，
	//其实内容是先写入到缓冲区中，所以需要调用Flush方法，将缓冲区的内容写入到文件中
	writer.Flush()

}
