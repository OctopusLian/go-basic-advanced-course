package main

import (
	"fmt"
	"time"
	"runtime"
)
// 全局定义channel， 用来完成数据同步
var channel = make(chan int)

// 定义一台打印机
func printer(s string)  {
	for _, ch := range s {
		fmt.Printf("%c", ch)			// 屏幕：stdout
		time.Sleep(300 * time.Millisecond)
	}
}

// 定义两个人使用打印机
func person1()  {				// person 先执行。
	printer("hello")
	channel <- 8
}
func person2()  {				// person 后执行
						// 		<- channel
	printer("wrold")
}

func main()  {
	go person1()
	<- channel

	go person2()

	runtime.Goexit()
}
