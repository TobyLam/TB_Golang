package main
import(
	"fmt"
)

func main(){

	//重点讲解 /、%
	//说明，如果参与运算的数都是整数，那么除后，去掉小数部分，保留整数部分
	fmt.Println(10 / 4)  // 2
	var n1 float32 = 10 /4 //
	fmt.Println(n1)

	//如果我们希望保留小数部分，则需要有浮点数参与计算
	var n2 = 10.0 /4
	fmt.Println(n2)

	// 演示 % 的使用
	// 看一个公式 a % b = a - a / b * b
	// 商（a/b）* 除数 (b) + 余数 (a%b) = 被除数（a）
	// 取余 = 被除数 - 商 * 除数
	fmt.Println("10%3=",10 % 3)      // = 1
	fmt.Println("-10%3=",-10 % 3)    // = -1
	fmt.Println("10%-3=",10 % -3)    // = 1
	fmt.Println("-10%-3=",-10 % -3)  // = -10 - （-10 / -3） * （-3） = -10 - 3 *（-3）= -10 - （-9）= -1
	fmt.Println("-5%3=",-5 % 3)


	// ++ 和 -- 的使用
	var i int = 10
	i++ // 等价 i = i + 1
	fmt.Println("i=",i) // 11
	i-- // 等价 i = i - 1
	fmt.Println("i=",i) // 10

}
