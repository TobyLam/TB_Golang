package main

import (
	"fmt"
	"go_code/familyAccount/utils"
)

func main() {
	fmt.Println("这个是面向对象的方式完成~~")
	//主界面
	utils.NewFamilyAccount().MainMenu()
}
