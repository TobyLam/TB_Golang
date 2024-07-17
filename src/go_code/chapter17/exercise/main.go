package main

import (
	"fmt"
	"reflect"
)

func main() {
	var str string = "tom" //ok
	//fs := reflect.ValueOf(str) //ok fs->string
	fs := reflect.ValueOf(&str) //ok fs->string
	//fs.SetString("jack")        //error
	fs.Elem().SetString("jack")
	fmt.Printf("%v\n", str)
}
