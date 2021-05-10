package main

import "fmt"

// map做函数参数、返回值，传引用
func mapDelete(m map[int]string, key int){
	delete(m, key)		// 删除 m 中 键值为 key的 map 元素
}

func main()  {

	m := map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}
	fmt.Println("before delete m :", m)

	mapDelete(m, 130)

	fmt.Println("after delete m :", m)
}
