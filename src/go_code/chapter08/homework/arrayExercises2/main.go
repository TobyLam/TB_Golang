package main

import (
    "fmt"
    "math"
    "math/rand"
    "time"
)

// 插入数据并排序
func insertAndSort(slice []float64, num float64) []float64 {
    slice = append(slice, num)
    //冒泡排序
    for i := 0; i < len(slice)-1; i++ {
        for j := 0; j < len(slice)-1-i; j++ {
            if slice[j] > slice[j+1] {
                slice[j], slice[j+1] = slice[j+1], slice[j]
            }
        }
    }
    return slice
}

func main() {
    /**
      已知有个排序好（升序）的数组，要求插入一个元素，最后打印该数组，顺序依然是升序
    */
    max := 100.0
    min := 0.0
    precision := 2

    var arr [10]float64
    rand.Seed(time.Now().Unix())

    for i := 0; i < len(arr); i++ {
        arr[i] = math.Floor(rand.Float64()*(max-min)*math.Pow10(precision)+min*math.Pow10(precision)) / math.Pow10(precision)
    }

    //冒泡排序
    for i := 0; i < len(arr); i++ {
        for j := 0; j < len(arr)-1-i; j++ {
            if arr[j] > arr[j+1] {
                arr[j], arr[j+1] = arr[j+1], arr[j]
            }
        }
    }

    fmt.Printf("原数组=%v\n", arr)

    num := 0.0
    fmt.Println("请输入需要插入的数字:")
    fmt.Scanln(&num)

    //声明一个切片
    var slice = arr[:]
    //插入元素并排序
    newSlice := insertAndSort(slice, num)
    fmt.Printf("插入元素后的结果：%v\n", newSlice)
}
