package main

import (
	"bufio"
	"fmt"
	_ "io"
	"os"
	"strconv"
	"strings"
)

type ValNode struct {
	row int
	col int
	val int
}

/**
 * 稀疏数组和常规数组的相互转换
 * 常用于压缩
 * 区别 sparsearray, 存盘【文件操作】
 */
func main() {

	//原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1
	chessMap[2][3] = 2

	//存盘，写入文件
	filePath := "f:/go_workspace/TB_Golang/chessmap.data"
	//打开文件
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("writeFile open file err=%v\n", err)
		return
	}

	//及时关闭file句柄
	defer file.Close()

	var sparseArr []ValNode

	valNode := ValNode{
		row: 11,
		col: 11,
		val: 0,
	}
	sparseArr = append(sparseArr, valNode)

	for i, v := range chessMap {
		for j, v2 := range v {
			if v2 != 0 {
				//创建一个ValNode 值节点
				valNode := ValNode{
					row: i,
					col: j,
					val: v2,
				}
				sparseArr = append(sparseArr, valNode)
			}
		}
	}

	//把稀疏数组写入文件
	writer := bufio.NewWriter(file)
	for _, valNode := range sparseArr {
		str := fmt.Sprintf("%d %d %d \n", valNode.row, valNode.col, valNode.val)
		writer.WriteString(str)
	}
	writer.Flush()

	file2, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Println("read open file err=", err)
		return
	}

	defer file2.Close()

	//读取文件内容，恢复原始数组
	reader := bufio.NewReader(file2)

	//循环读取文件
	var index = 0 //读取到的数据行号，从0开始
	var chessMap2 [][]int

	for {

		line, err := reader.ReadBytes('\n') //返回读到数据和\n字节切片
		if err != nil {
			break
		}

		index++ //行号递增
		//将每行数据使用空格分割
		tmp := strings.Split(string(line), " ")
		row, _ := strconv.Atoi(tmp[0])
		col, _ := strconv.Atoi(tmp[1])
		val, _ := strconv.Atoi(tmp[2])

		if index == 1 {
			//第一行，表示原数组的行、列数，默认值

			//构建row行，col列的切片
			for i := 0; i < row; i++ {
				var tmp_slice []int
				for j := 0; j < col; j++ {
					tmp_slice = append(tmp_slice, val)
				}
				chessMap2 = append(chessMap2, tmp_slice)
			}

		} else {
			chessMap2[row][col] = val
		}

	}
	fmt.Println("读取文件结束")
	// 打印数据

	fmt.Println("从磁盘读取后的数据")
	for _, v := range chessMap2 {
		for _, v1 := range v {
			fmt.Printf("%d\t", v1)
		}
		fmt.Println()
	}
}
