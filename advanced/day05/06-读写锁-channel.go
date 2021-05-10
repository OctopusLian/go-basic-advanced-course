package main

import (
	"math/rand"
	"time"
	"fmt"
)

//var value06 int		// 定义全局变量，模拟共享数据
func readGo06(in <-chan int, idx int)  {
	for {
		num := <-in				// 从 channel 中读取数据
		fmt.Printf("----%dth 读 go程，读出：%d\n", idx, num)
		time.Sleep(time.Second)
	}
}
func writeGo06(out chan<- int, idx int)  {
	for {
		// 生成随机数
		num := rand.Intn(1000)
		out <- num					// 写入channel
		fmt.Printf("%dth 写go程，写入：%d\n", idx, num)
		time.Sleep(time.Millisecond * 300)		// 放大实验现象
	}
}
func main()  {
	// 播种随机数种子
	rand.Seed(time.Now().UnixNano())
	ch := make(chan int)

	for i:=0; i<5; i++ {			// 5 个 读 go 程
		go readGo06(ch, i+1)
	}
	for i:=0; i<5; i++ {			//
		go writeGo06(ch, i+1)
	}
	for{
		;
	}
}
