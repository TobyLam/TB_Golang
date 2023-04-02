package main
import(
	"fmt"
)

func main(){

	//在golang中, ++ 和 -- 只能独立使用.
	var i int = 8
	var a int
	// a = i++ //错误，i++只能独立使用
	// a = i-- //错误，i--只能独立使用

	//if i++ > 0 { //错误，i++只能独立使用
	//	fmt.Println("ok")
	//}

	i++
	//++i //错误，在golang没有前 ++
	i--
	//--i //错误，在golang没有前 --

	a = i
	fmt.Println(a)
}
