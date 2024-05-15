package main
import(
    "fmt"
)

func printArr(arr *[3][4]float64){
    for i:= 0;i<len(*arr);i++{
        for j:= 0;j<len((*arr)[i]);j++{
           fmt.Printf("%v\t",(*arr)[i][j])
        }
        fmt.Println()
    }
}

func main(){
    /**
    * 定义一个3行4列的二维数组，逐个从键盘输入值，编写程序将四周的数据清0
    * 四周清0: 第一行和最后一行均为0, 中间行的第一个元素和最后一个元素为0
    */
    var arr [3][4]float64
    //赋值前遍历
    fmt.Println("原始数组:")
    printArr(&arr)
    //赋值
    for i:=0;i<len(arr);i++{
        for j:=0;j<len(arr[i]);j++{
            fmt.Printf("请输入数组%v行%v列的元素:\n",i+1,j+1)
            fmt.Scanln(&arr[i][j])
        }
    }
    //复制后遍历
    fmt.Println("赋值后数组")
    printArr(&arr)
    //四周清0
    for i:= 0; i<len(arr);i++{
       for j:=0;j<len(arr[i]);j++{
           //以下逻辑可放在输入赋值时进行
           if i== 0 || i== len(arr)-1{
               arr[i][j] = 0
           }else{
               if j == 0 || j== len(arr[i])-1{
                   arr[i][j] = 0
               }
           }
       }
    }
    //清0后遍历
    fmt.Println("清0后数组:")
    printArr(&arr)
    


}
