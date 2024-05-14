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
//二分查找
func binaryFind(arr *[10]int,leftIndex int,rightIndex int,findVal int){

   if leftIndex > rightIndex {
     fmt.Printf("%v没找到\n",findVal)
     return
   }
   //计算中间的下标
   middle := (leftIndex + rightIndex)/2
   
   if (*arr)[middle] > findVal {
      //中间值大于查找值,则查找值可能在上半区间 leftIndex --> middle-1
      binaryFind(arr,leftIndex,middle-1,findVal)
   }else if (*arr)[middle] < findVal {
      //中间值小于查找值，则查找值可能在下半区间 middle+1 --> rightIndex
      binaryFind(arr,middle+1,rightIndex,findVal)
   }else{
      //找到了
      fmt.Printf("%v找到了,下标为%v\n",findVal,middle)
      return
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
    //求平均值、最大值、最大值的下标
    sum := 0
    max := 0
    maxIndex := 0
    for i,v := range arr2 {
       sum += v
       if v >= max {
           max = v
           maxIndex = i
       }
    }
    fmt.Printf("平均值%v 最大值%v 最大值的下标%v\n",sum/len(arr2), max, maxIndex)
    //是否有55
    binaryFind(&arr2,0,len(arr2)-1,55)
    
    fmt.Println(arr2)
}
