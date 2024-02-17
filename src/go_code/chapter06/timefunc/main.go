package main

import (
	"fmt"
	"time"
)

func main() {
	//看看日期和时间相关函数和使用方法
	//1、获取当前时间
	now := time.Now()
	fmt.Printf("now=%v now type=%T\n", now, now)

	//2.通过now获取年月日时分秒
	fmt.Printf("年=%v\n", now.Year())
	fmt.Printf("月=%v\n", int(now.Month()))
	fmt.Printf("日=%v\n", now.Day())
	fmt.Printf("时=%v\n", now.Hour())
	fmt.Printf("分=%v\n", now.Minute())
	fmt.Printf("秒=%v\n", now.Second())

	//格式化日期时间
	fmt.Printf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d \n", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	dateStr := fmt.Sprintf("当前年月日 %02d-%02d-%02d %02d:%02d:%02d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	fmt.Printf("dateStr=%v\n", dateStr)

	//格式化日期时间的第二种方式
	fmt.Printf(now.Format("2006/01/02 15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("2006-01-02"))
	fmt.Println()
	fmt.Printf(now.Format("15:04:05"))
	fmt.Println()
	fmt.Printf(now.Format("06"))
	fmt.Println()
	fmt.Printf(now.Format("05"))
	fmt.Println()

	//需求，每隔1秒钟打印一个数字，打印到100时就退出
	//需求2，每隔0.1秒钟打印一个数字，打印到100时就退出
	i := 0
	for {
		i++
		fmt.Println(i)
		//休眠1秒
		//time.Sleep(time.Second)
		//休眠0.1秒
		time.Sleep(time.Millisecond * 100)
		if i == 10 {
			break
		}
	}

	//Unix 和 UnixNano 的使用
	fmt.Printf("unix时间戳=%v unixnano时间戳=%v\n", now.Unix(), now.UnixNano())
}
