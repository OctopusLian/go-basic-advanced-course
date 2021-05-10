package main

import (
	"time"
	"fmt"
)

/*func main()  {
	fmt.Println("当前时间：", time.Now())
	// 创建定时器
	myTimer := time.NewTimer(time.Second * 2)
	nowTime := <- myTimer.C  // chan 类型
	fmt.Println("现下时间：", nowTime)
}*/

// 3 种定时方法
/*func main()  {
	// 1 . sleep
	time.Sleep(time.Second)

	// 2. Timer.C
	myTimer := time.NewTimer(time.Second * 2)		// 创建定时器， 指定定时时长
	nowTime := <- myTimer.C  						// 定时满，系统自动写入系统时间
	fmt.Println("现下时间：", nowTime)

	// 3 time.After
	fmt.Println("当前时间：", time.Now())
	nowTime2 := <-time.After(time.Second * 2)
	fmt.Println("现下时间：", nowTime2)
}*/

// 定时器的停止和重置
func main()  {

	myTimer := time.NewTimer(time.Second * 10)		// 创建定时器。
	myTimer.Reset(1 * time.Second)		// 重置定时时长为 1
	go func() {
		for {
			<- myTimer.C
			fmt.Println("子go程，定时完毕")
		}
	}()

	//myTimer.Stop()		// 设置定时器停止
	for {
		;
	}
}

