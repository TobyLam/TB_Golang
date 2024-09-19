package main

import (
	"fmt"
)

// 编写一个函数，完成找路
// myMap : 地图，同一个地图，使用引用
// i,j表示对地图的哪个点进行测试
func SetWay(myMap *[8][7]int, i, j int) bool {

	//分析什么情况，就找到出路
	//myMap[6][5]==2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明要继续找
		if myMap[i][j] == 0 { //如果这个点是可以探测的

			//假设这个点可以通,但是需要探测,上下左右[运行结果有问题]
			//更换策略 下右上左
			myMap[i][j] = 2
			// else... if ... 只能往一个方向
			if SetWay(myMap, i+1, j) { //向下走
				return true
			} else if SetWay(myMap, i, j+1) { //向右走
				return true
			} else if SetWay(myMap, i-1, j) { //向上走
				return true
			} else if SetWay(myMap, i, j-1) { //往左走
				return true
			} else { //当前点死路
				myMap[i][j] = 3
				return false
			}

		} else { //说明这个点不能探测，为1，是墙
			return false
		}
	}

}

func main() {
	//先创建一个二维数组，模拟迷宫
	//规则
	//1.如果元素的值为1，就是墙
	//2.如果元素的值为0，是没有走过的点
	//3.如果元素的值为2，是一个通路
	//4.如果元素的值为3，是走过的点，但是走不通
	var myMap [8][7]int

	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	//把地图的最左和最右设置为1
	for i := 0; i < 8; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	//加墙
	myMap[3][1] = 1
	myMap[3][2] = 1

	//堵死
	myMap[1][2] = 1
	myMap[2][2] = 1

	//打印地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

	//测试
	SetWay(&myMap, 1, 1)
	fmt.Println("探测完毕..")

	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
