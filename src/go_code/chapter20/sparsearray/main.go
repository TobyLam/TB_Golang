package main

import (
	"fmt"
)

type ValNode struct {
	row int //行
	col int //列
	//val interface{} //不确定类型，使用空接口类型
	val int
}

func main() {

	//1. 先创建一个原始数组
	var chessMap [11][11]int
	chessMap[1][2] = 1 //黑子
	chessMap[2][3] = 2 //蓝子

	//输出原始数组
	for _, v := range chessMap {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}

	//3.转成稀疏数组
	// 思路
	//(1).遍历 chessMap,如果发现一个元素的值不为0，创建一个node结构体
	//(2).将其放入到对应的切片中

	var sparseArr []ValNode

	//标准的一个稀疏数组有一个 记录元素的二维数组的规模（行和列，默认值）
	//创建一个ValNode 值节点
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

	//输出稀疏数组
	fmt.Println("当前的稀疏数组是::::")
	for i, valNode := range sparseArr {
		fmt.Printf("%d: %d %d %d\n", i, valNode.row, valNode.col, valNode.val)
	}

	//将这个稀疏数组，存盘 f:/go_workspace/chessmap.data

	//如何恢复原始的数组

	//1.打开这个f:/go_workspace/chessmap.data => 恢复原始数组。

	//2. 这里使用稀疏数组恢复
	var chessMap2 [11][11]int //可以使用切片，读取原始数组的行、列数、默认值

	// 遍历 sparseArr [遍历文件每一行]
	for i, valNode := range sparseArr {
		if i != 0 { //跳过第一行记录值
			chessMap2[valNode.row][valNode.col] = valNode.val
		}

	}

	//查看chessMap2 是否已恢复
	fmt.Println("恢复后的原始数据.....")
	for _, v := range chessMap2 {
		for _, v2 := range v {
			fmt.Printf("%d\t", v2)
		}
		fmt.Println()
	}
}
