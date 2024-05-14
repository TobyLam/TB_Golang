package main
import(
    "fmt"
    "math/rand"
    "time"
)

//冒泡排序
func bubbleSort(arr *[10]int){
 
    for i := 0;i < len(*arr)-1;i++{
    	for j:= 0;j < len(*arr)-i-1;j++{

	    if(*arr)[j] < (*arr)[j+1] {
               (*arr)[j],(*arr)[j+1] = (*arr)[j+1],(*arr)[j]	        
            }

	}
    }
}

func main(){
    /*
    随机生成10个整数(1_100的范围)保存到数组，并倒序打印
    以及求平均值、求最大值和最大值的下标、并查找里面是否有55
    */

    //声明数组
    var arr [10]int
    //生成随机种子    
    rand.Seed(time.Now().Unix())
    //生成随机整数并保存到数组
    for i := 0;i<len(arr);i++{
	arr[i] = rand.Intn(100)+1	
    }
    arr2 := arr
    //排序
    bubbleSort(&arr)
    //倒序打印 	
    fmt.Println(arr)
    //求平均值、最大值、最大值的下标，是否有55
    sum := 0
    max := 0
    maxIndex := 0
    existFF := 0
    for i,v := range arr2 {
       sum += v
       if v >= max {
           max = v
           maxIndex = i
       }
       if v == 55 {
           existFF = 1
       }    
    }
    fmt.Printf("平均值%v 最大值%v 最大值的下标%v\n",sum/len(arr2), max, maxIndex)
    if existFF == 1 {
       fmt.Println("有55")
    }else{
       fmt.Println("没有55")
    }
    fmt.Println(arr2)
}
