package main

import (
	"fmt"
)

func sing()  {
	for i:=0;i<50;i++ {
		fmt.Println("----正在唱：隔壁泰山----")
		//time.Sleep(100 *time.Millisecond)
	}
}

func dance()  {
	for i:=0;i<50;i++ {
		fmt.Println("----正在跳舞：赵四街舞----")
		//time.Sleep(100 * time.Millisecond)
	}
}

func main()  {
	go sing()
	go dance()

	for {
		;
	}
}
