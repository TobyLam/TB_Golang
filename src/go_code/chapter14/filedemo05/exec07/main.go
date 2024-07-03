package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// 定义一个结构体，用于保存统计结果
type CharCount struct {
	ChCount    int //记录英文字母个数
	NumCount   int //记录数字个数
	SpaceCount int //记录空格个数
	OtherCount int //记录其他字符个数
}

func main() {
	//思路：打开一个文件，创建一个Reader
	//每读取一行，就去统计改行有多少个 英文、数字、空格 和其他字符
	//然后将结构保存到一个结构体

	fileName := "f:/go_workspace/TB_Golang/hello.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("open file err:", err)
		return
	}

	//关闭文件
	defer file.Close()

	//定义一个CharCount 结构体实例
	var count CharCount
	//创建一个Reader
	reader := bufio.NewReader(file)

	//循环读取fileName的内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		//兼容中文
		str2 := []rune(str)

		//遍历 str,进行统计
		for _, v := range str2 {
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough //穿透
			case v >= 'A' && v <= 'Z':
				count.ChCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ':
				count.SpaceCount++
			default:
				count.OtherCount++

			}
		}
	}

	//输出统计的结果
	fmt.Printf("英文字母个数=%d,数字个数=%d,空格个数=%d,其他字符个数=%d\n", count.ChCount,
		count.NumCount, count.SpaceCount, count.OtherCount)

}
