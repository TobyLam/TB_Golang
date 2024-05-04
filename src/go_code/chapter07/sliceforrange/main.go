package main

import (
	"fmt"
)

func main() {
	//使用常规的for循环遍历切片
	var arr [5]int = [...]int{10, 20, 30, 40, 50}
	//slice := arr[1:4] // 20,30,40
	//slice := arr[0:len(arr)]
	slice := arr[:]
	for i := 0; i < len(slice); i++ {
		fmt.Printf("slice[%v]=%v ", i, slice[i])
	}

	fmt.Println()
	//使用for--range 方式遍历切片
	for i, v := range slice {
		fmt.Printf("i=%v v=%v\n", i, v)
	}

	//fmt.Println(slice[3])
	fmt.Println()
	slice2 := slice[2:4]

	slice2[0] = 100 // 因为arr,slice和slice2 指向的数据空间是同一个，因此slice2[0]=100,其它的都变化

	fmt.Println("slice2=", slice2)
	fmt.Println("slice=", slice)
	fmt.Println("arr=", arr)

	fmt.Println()

	//用 append 内置函数，可以对切片进行动态追加
	var slice3 []int = []int{100, 200, 300}
	//通过append直接给slice3追加具体的元素
	slice3 = append(slice3, 400, 500, 600)
	//slice4 := append(slice3, 400, 500, 600)
	fmt.Println("slice3=", slice3)
	//通过append将切片slice3追加给slice3
	slice3 = append(slice3, slice3...)

	fmt.Println("slice3=", slice3)

	var sliceA []int = make([]int, 5, 10)
	fmt.Println("sliceA=", sliceA)
	fmt.Println("sliceA[0]的地址", &sliceA[0])
	sliceA1 := append(sliceA, 111, 222, 333)
	fmt.Println("sliceA=", sliceA)
	fmt.Println("sliceA[0]的地址", &sliceA[0])
	fmt.Println("sliceA1=", sliceA1)
	fmt.Println("sliceA1[0]的地址", &sliceA1[0])

	// append后，原切片的容量不够，地址发生变化
	sliceA3 := append(sliceA, 11, 22, 33, 44, 55, 66)
	fmt.Println("sliceA3=", sliceA3)
	fmt.Println("sliceA3[0]的地址", &sliceA3[0])

	arr2sliceA2 := [5]int{1, 2, 3, 4, 5}
	var sliceA2 []int = arr2sliceA2[:]
	fmt.Println("sliceA2=", sliceA2)
	fmt.Println("cliceA2容量", cap(sliceA2))

	//切片的拷贝操作
	//切片使用 copy 内置函数完成拷贝，举例说明
	var slice4 []int = []int{1, 2, 3, 4, 5}
	var slice5 = make([]int, 10)
	copy(slice5, slice4)
	fmt.Println("slice4=", slice4) // 1,2,3,4,5
	fmt.Println("slice5=", slice5) // 1,2,3,4,5,0,0,0,0，0

	var a []int = []int{1, 2, 3, 4, 5}
	var slice2a = make([]int, 1)
	fmt.Println(slice2a) // [0]
	copy(slice2a, a)     //不会报错
	fmt.Println(slice2a) // [1]

}
