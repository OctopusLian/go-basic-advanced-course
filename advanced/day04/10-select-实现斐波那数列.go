package main

import (
	"fmt"
	"runtime"
)

func fibonacci(ch <-chan int, quit <-chan bool)  {
	for {
		select {
		case num := <-ch :
			fmt.Print(num , " ")
		case <-quit:
			//return
			runtime.Goexit()		//等效于 return
		}
	}
}

func main()  {
	ch := make(chan int)
	quit := make(chan bool)

	go fibonacci(ch, quit)		// 子go 程 打印fibonacci数列

	x, y := 1, 1
	for i:=0; i<40;i++ {
		ch <- x
		x, y = y, x+y
	}
	quit <- true
}
