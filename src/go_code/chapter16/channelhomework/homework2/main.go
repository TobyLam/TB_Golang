package main

import (
	"bufio"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
*
写入字符串到文件中
*/
func writeFile(name string, str string) {
	f, _ := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	writer := bufio.NewWriter(f)
	_, err := writer.WriteString(str)
	if err != nil {
		return
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

func randNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000)
}

func writeDataToFile(exitChan chan bool, index int) {
	var str string
	//for i := 0; i < 1000; i++ {
	for i := 0; i < 10; i++ {
		//v := rand.Int()
		v := randNum()

		str += strconv.Itoa(v) + "\n"
	}
	writeFile("src"+strconv.Itoa(index)+".txt", str)
	exitChan <- true
}

/*
*
冒泡排序
*/
func bubbleSort(arr *[10]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				temp := (*arr)[j]
				(*arr)[j] = (*arr)[j+1]
				(*arr)[j+1] = temp
			}
		}
	}
}

func sort(exitChan chan bool, index int) {
	f, _ := os.OpenFile("src"+strconv.Itoa(index)+".txt", os.O_RDONLY, 0666)
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)
	reader := bufio.NewReader(f)
	//var arr [1000]int
	var arr [10]int
	arrIndex := 0
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		str = strings.TrimSuffix(str, "\n") //返回去除str可能的后缀suffix的字符串。
		rs, _ := strconv.Atoi(str)
		arr[arrIndex] = rs
		arrIndex++
	}
	bubbleSort(&arr)
	var str2 string
	for i := 0; i < len(arr); i++ {
		str2 += strconv.Itoa(arr[i]) + "\n"
	}
	writeFile("dest"+strconv.Itoa(index)+".txt", str2)
	exitChan <- true
}

func main() {
	exitChan := make(chan bool, 10)
	for i := 0; i < 10; i++ {
		go writeDataToFile(exitChan, i)
	}
	for i := 0; i < 10; i++ {
		<-exitChan
	}
	for i := 0; i < 10; i++ {
		go sort(exitChan, i)
	}
	for i := 0; i < 10; i++ {
		<-exitChan
	}
}
