package main
import(
    "fmt"
)

func arrayForRange(arr *[4][4]int){
    for i := 0;i<len(*arr);i++{
       for j:=0;j<len((*arr)[i]);j++{
           fmt.Printf("%v\t",(*arr)[i][j])
       }
       fmt.Println()
    }
}

func main(){
   /**
   * 定义一个4行4列的二维数组,逐个从键盘输入值,然后将第1行和第4行的数据进行交换，
   * 将第2行和第3行的数据进行交换
   */
   arr := [4][4]int{}
   fmt.Println("初始化数组:")
   arrayForRange(&arr)

   //遍历从键盘输入值
   for i := 0;i<len(arr);i++{
      for j:= 0;j<len(arr[i]);j++{
          fmt.Printf("请输入第%v行第%v列的数据:",i+1,j+1)
          fmt.Scanln(&arr[i][j])
      }
   }
   //输入后的数组
   fmt.Println("赋值后数组:")
   arrayForRange(&arr)
   //数据交换
   arr[0],arr[3] = arr[3],arr[0]
   arr[1],arr[2] = arr[2],arr[1]
   
   fmt.Println("交换后数组:")
   arrayForRange(&arr)
 
}
