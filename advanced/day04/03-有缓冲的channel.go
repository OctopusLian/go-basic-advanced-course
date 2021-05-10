package main

import (
	"fmt"
)

func main()  {
	ch := make(chan int, 5)		// 存满3个元素之前，不会阻塞
	fmt.Println("len=", len(ch), "cap=", cap(ch))

	go func() {
		for i:=0; i< 10;i++ {
			ch <- i
			len := len(ch)
			cap := cap(ch)
			fmt.Println("子go程：i", i, "len=", len, "cap=", cap)
			//fmt.Println("子go程：i", i)
		}
	}()
	//time.Sleep(time.Second * 3)
	for i:=0; i< 8;i++ {
		num := <- ch
		fmt.Println("主go程读到：", num)
	}
}
