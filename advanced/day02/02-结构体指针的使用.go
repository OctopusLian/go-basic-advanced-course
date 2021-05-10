package main

import (
	"fmt"
	"unsafe"
)

type Person2 struct {
	name string
	sex byte
	age int
}


func test2(p *Person2)  {
	fmt.Println("test2:", unsafe.Sizeof(p))
	p.name = "Luffy"
	p.age = 779
	p.sex = 'm'
}
func main()  {
	var p1 *Person2 = &Person2{"n1", 'f', 19}
	fmt.Println("p1", p1)

	var tmp Person2 = Person2{"n1", 'f', 19}
	var p2 *Person2
	p2 = &tmp
	fmt.Println("p2", p2)

	p3 := new(Person2)
	p3.name = "n3"
	p3.age = 22
	p3.sex = 'f'
	fmt.Printf("p3, type= %T\n", p3)
	fmt.Println("p3:", p3)

	fmt.Println("main:", unsafe.Sizeof(p3))

	test2(p3)

	fmt.Println("p3:", p3)
}
