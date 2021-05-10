package main

import (
	"fmt"
	"time"
	"runtime"
)


func main()  {
	ch := make(chan int)			// 用来进行数据通信的 channel
	quit := make(chan bool)			// 用来判断是否退出的 channel
	//ch2 := make(chan string)
	go func() {					// 写数据
		for i:=0; i<5;i++{
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
		quit <- true			// 通知主go程 退出
		runtime.Goexit()
	}()
	for {						// 主go程 select 监听 chan数据流动
		select {
		case num := <-ch:		// 不可读，阻塞。可以读，将数据保存至num
			fmt.Println("读到：", num)		// 模拟使用数据

		case <-quit:			// 不可读，阻塞。可以读，将主go程结束。
			break 				// break 跳出 select	不可用
			//runtime.Goexit()	// 终止 主 go 程		不可用
			//return 				// 终止进程
		}
		fmt.Println("============")		// select 自身不带有循环机制，需借助外层 for 来循环监听
	}
}
