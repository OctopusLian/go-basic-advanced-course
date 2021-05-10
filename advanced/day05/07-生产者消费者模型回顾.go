package main

import (
	"fmt"
	"time"
	"math/rand"
)

func producer(out chan<- int, idx int)  {
	for i:=0; i<50; i++ {
		num := rand.Intn(800)
		fmt.Printf("生产者%dth，生产：%d\n", idx, num)
		out <- num
	}
	close(out)
}

func consumer(in <-chan int, idx int)  {
	for num := range in {
		fmt.Printf("-----消费者%dth，消费：%d\n",idx,  num)
	}
}

func main()  {
	product := make(chan int)
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<5; i++ {
		go producer(product, i+1)			// 1 生产者
	}
	for i:=0; i<5; i++ {
		go consumer(product, i+1)			// 3 个消费者
	}
	for {
		;
	}
}
