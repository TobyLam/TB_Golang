package main
import(
	"fmt"
)

func main(){

	var i int = 5
	//二进制输出
	fmt.Printf("%b",i)

	//八进制：0-7 ，满 8 进 1. 以数字 0 开头表示。
	var j int = 011 // 011 =>
	fmt.Println("j=", j)

	//十六进制：0-9 及 A-F，满 16 进 1. 以 0x 或 0X 开头表示
	var k int = 0x11 // 0x11 => 17
	fmt.Println("k=",k)
}
