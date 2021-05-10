package main

import "fmt"

type Person3 struct {
	name string
	age int
	flg bool
	intereset []string
}

// 通过函数参数 初始化 结构体
func initFunc(p *Person3)  {
	p.name = "Nami"
	p.age = 18
	p.flg = true
	p.intereset = append(p.intereset, "shopping")
	p.intereset = append(p.intereset, "sleeping")
}

// 通过函数返回值，初始化结构体。
func initFunc2() *Person3 {
	p := new(Person3)
	p.name = "Nami"
	p.age = 18
	p.flg = true
	p.intereset = append(p.intereset, "shopping")
	p.intereset = append(p.intereset, "sleeping")
	return p		// 返回指针变量的值 —— heap的地址。
}

func main()  {
/*	var person Person3
	initFunc(&person)
	fmt.Println("person:", person)*/

	p2 := initFunc2()
	fmt.Println("p2", p2)

}
