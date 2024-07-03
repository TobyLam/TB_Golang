package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { //文件或目录存在
		return true, nil
	}
	if os.IsNotExist(err) { //文件或目录不存在
		return false, nil
	}
	return false, err
}

func main() {
	//判断文件是否存在
	filePath := "f:/go_workspace/TB_Golang/hellod.txt"

	res, err := PathExists(filePath)

	fmt.Println(res, err)
}
