package main

import (
	"fmt"
	"math"
)

func extremeValue(arr *[8]float64) {
	max := arr[0]
	min := arr[0]
	for i := 0; i < len(*arr); i++ {
		if max <= (*arr)[i] {
			max = (*arr)[i]
		}
		if min >= (*arr)[i] {
			min = (*arr)[i]
		}
	}

	for i := 0; i < len(*arr); i++ {
		if (*arr)[i] == max {
			fmt.Printf("评委%v打了最高分%v\n", i+1, max)
		}
		if (*arr)[i] == min {
			fmt.Printf("评委%v打了最低分%v\n", i+1, min)
		}
	}
}

// 排序
func scoresSort(arr [8]float64) [8]float64 {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	sortArr := arr
	return sortArr
}

// 找到最佳评委和最差评委
func findExtremeJudge(arr [8]float64, score float64) {
	minAbs := math.Abs(arr[0] - score)
	maxAbs := math.Abs(arr[0] - score)

	for i := 0; i < len(arr); i++ {
		//fmt.Printf("评委%v的打分和最后得分相差%v\n", i+1, math.Abs(arr[i]-score))
		if math.Abs(arr[i]-score) < minAbs {
			minAbs = math.Abs(arr[i] - score)
		}
		if math.Abs(arr[i]-score) > maxAbs {
			maxAbs = math.Abs(arr[i] - score)
		}
	}
	var bestJudgeIndexSlice []int = make([]int, 1)
	var worstJudgeIndexSlice []int = make([]int, 1)

	for i := 0; i < len(arr); i++ {
		//fmt.Printf("评委%v的打分和最后得分相差%v\n", i+1, math.Abs(arr[i]-score))
		if math.Abs(arr[i]-score) == minAbs {
			bestJudgeIndexSlice = append(bestJudgeIndexSlice, i+1)
		}
		if math.Abs(arr[i]-score) == maxAbs {
			worstJudgeIndexSlice = append(worstJudgeIndexSlice, i+1)
		}
	}

	fmt.Printf("最佳评委是评委%v,最差评委是评委%v\n", bestJudgeIndexSlice[1:], worstJudgeIndexSlice[1:])
}

func main() {

	/**
	 * 跳水比赛8个评委打分。运动员的成绩是8个成绩取掉一个最高分，去掉一个最低分，剩
	 * 下的6个分数的平均分就是最后得分。使用一维数组实现如下功能:
	 * (1) 请把打最高分的评委和最低分的评委找出来。
	 * (2) 找出最佳评委和最差评委。最佳评委就是打分和最后得分最接近的评委。最差评委就是打分和最后得分相差最大的。
	 */

	//定义一个一维数组，保存8个评委的打分
	scores := [8]float64{}

	for i := 0; i < len(scores); i++ {
		fmt.Printf("请第%v个评委打分：", i+1)
		fmt.Scanln(&scores[i])
	}

	fmt.Println("评委打分结果如下:", scores)
	// 请把打最高分的评委和最低分的评委找出来
	extremeValue(&scores)

	// 排序，并返回到一个新的数组
	sortScores := scoresSort(scores)
	fmt.Println("排序后的评委打分结果如下:", sortScores)

	//切片
	scoreSlice := sortScores[1:7]
	fmt.Println("去掉一个最高分，去掉一个最低分后的打分结果如下:", scoreSlice)

	// 计算最终得分
	sumScore := 0.0
	for _, v := range scoreSlice {
		sumScore += v
	}
	lastScore := sumScore / float64(len(scoreSlice))
	fmt.Printf("运动员最后得分：%v\n", lastScore)

	//找出最佳评委和最差评委
	findExtremeJudge(scores, lastScore)

}
