package main

import (
	"fmt"
)

func main(){
	var num int = 9
	fmt.Printf("num address=%v \n",&num)

	var ptr *int
	ptr = &num
	*ptr = 10 //这里修改时，会到num的值变化
	fmt.Println("num =", num)

	is_zero := true

	var ptr2 *bool
	ptr2 = &is_zero

	fmt.Printf("ptr2 address=%v \n",ptr2)
	fmt.Printf("ptr2指向的值=%v",*ptr2)
}
