package main

import "fmt"

// 快速排序
// 说明
// 1.left表示 数组左边的下标
// 2.right表示 数组右边的下标
// 3.array 表示要排序的数组
func QuickSort(left int, right int, array *[10]int) {
	//fmt.Println("__________________________")
	l := left
	r := right

	//fmt.Printf("初始左下标%d,右下标%d\n", l, r)

	// pivot 中轴，以该数作为中轴(支点)，左右分割
	pivot := array[(left+right)/2] //可以不使用中轴，但中轴的效率较快

	//fmt.Printf("pivot=%d\n", pivot)

	//for 从小到大排序，所以循环的目标是：将比pivot 小的数放到左边，比pivot 大的数放到右边
	for l < r {
		//fmt.Println("再进循环时l=", l)
		//从 pivot 的左边找到大于等于pivot的值
		for array[l] < pivot { //左边的数据小于中轴的数据，继续查找，直到找到大于等于中轴数据的数，获得下标
			l++ //从左往右找的，自增
			//fmt.Println("寻找过程l=", l)
		}

		//fmt.Println("再进循环时r=", r)
		//从 pivot 的右边找到小于等于pivot的值
		for array[r] > pivot { //右边的数据大于中轴的数据，继续查找，直到找到小于等于中轴数据的数，获得下标
			r-- //从右往左找的，自减
			//fmt.Println("寻找过程r=", r)
		}

		// 表明本次分解任务完成，按照支点，分成大于支点的数为一堆，小于支点的数为一堆
		if l >= r {
			//fmt.Printf("触发一次。。。l=%d,r=%d", l, r)
			//fmt.Println(array)
			break
		}

		//交换左右的值
		//fmt.Printf("arr[%d]=%d与arr[%d]=%d进行交换\n", l, array[l], r, array[r])
		array[l], array[r] = array[r], array[l]
		fmt.Println(*array)
		fmt.Println()
		//fmt.Println("此时的数组", array)
		//fmt.Println()

		//优化
		//左右交换后，如果交换后的较小值(array[l])等于支点，则证明找到的满足条件的r下标的值等于支点，则r下标往后(>=r)的下标的值必定不满足条件，不需要再交换，r下标往前一位开始寻找满足条件（小于等于支点）的下标
		if array[l] == pivot {
			r--
		}
		//左右交换后，如果交换后的较大值(array[r])等于支点，则证明找到的满足条件的l下标的值等于支点，则l下标往前(>=l)的下标的值必定不满足条件，不需要再交换，l下标往后一位开始寻找满足条件（大于等于支点）的下标
		if array[r] == pivot {
			l++
		}
	}

	//如果 l==r,再移动一次,避免死循环
	//fmt.Println("退出循环后l=", l)
	//fmt.Println("退出循环后r=", r)
	if l == r {
		l++
		r--
	}

	//交换一次后，支点左右两侧的数据可能还是乱序，需要对子数组重复排序，直至左游标和右游标重叠，排序完成
	//向左递归
	if left < r {
		//fmt.Println("进行了一次左递归,left=", left, ",r=", r)
		QuickSort(left, r, array)
	}
	//向右递归
	if right > l {
		//fmt.Println("进行了一次右递归,right=", right, ",l=", l)
		QuickSort(l, right, array)
	}

}

func main() {

	arr := [10]int{-9, 78, 0, 23, -567, 70, 123, 90, 0, -23}
	fmt.Println("初始", arr)
	fmt.Println()

	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("main...")
	fmt.Println(arr)
}
