package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 编写一个函数，接收两个文件路径srcFileName dstFileName
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", srcFileName, err)
	}

	defer srcFile.Close()

	//通过srcFile,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file %s failed, err:%v\n", dstFileName, err)
	}

	//通过dstFile,获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	return io.Copy(writer, reader)
}
func main() {
	//将 c:/Users/TobyLam/Pictures/小马清清.jpg 文件拷贝到 f:/go_workspace/TB_Golang/maqq.jpg

	//调用CopyFile 完成文件拷贝
	srcFile := "c:/Users/TobyLam/Pictures/小马清清.jpg"
	dstFile := "f:/go_workspace/TB_Golang/maqq.jpg"
	_, err := CopyFile(dstFile, srcFile)
	if err == nil {
		fmt.Println("文件拷贝成功")
	} else {
		fmt.Printf("文件拷贝失败，err:%v\n", err)
	}

}
