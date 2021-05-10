package main

import "fmt"

func main()  {

/*	var m1 map[int]string			// 声明map ，没有空间，不能直接存储key -- value
	//m1[100] = "Green"
	if m1 == nil {
		fmt.Println("map is nil ")
	}

	m2 :=  map[int]string{}			//
	fmt.Println(len(m2))
	fmt.Println("m2 = ", m2)
	m2[4] = "red"
	fmt.Println("m2 = ", m2)

	m3 := make(map[int]string)
	fmt.Println(len(m3))
	fmt.Println("m3 = ", m3)
	m3[400] = "red"
	fmt.Println("m3 = ", m3)

	m4 := make(map[int]string, 5)		// len
	fmt.Println("len(m4) = ", len(m4))
	//fmt.Println("len(m4) = ", cap(m4))		// 不能在map中使用 cap（）
	fmt.Println("m4 = ", m4)*/

/*	// 初始化map
	var m5 map[int]string = map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}

	fmt.Println("m5 = ", m5)

	m6 := map[int]string{1:"Luffy", 130:"Sanji", 1303:"Zoro"}
	fmt.Println("m6 = ", m6)
*/

/*	m7 := make(map[int]string, 1)
	m7[100] = "Nami"
	m7[20] = "Hello"
	m7[3] = "world"
	fmt.Println("m7=", m7)

	m7[3] = "yellow"			// 成功！ 将原map中 key 值为 3 的map元素，替换。
	fmt.Println("m7=", m7)*/

	// 遍历map
/*	var m8 map[int]string = map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}
	for k, v := range m8 {
		fmt.Printf("key:%d --- value:%q\n", k, v)
	}

	// range返回的key/ value 。 省略value打印。
	for _, K := range m8 {
		fmt.Printf("key:%s\n", K)
	}*/

	// 判断 map 中的key 是否存在
	var m9 map[int]string = map[int]string{1:"Luffy", 130:"Sanji", 1301:"Zoro"}

	if v, has := m9[12]; has {	// m9[下标] 返回两个值，第一个是value，第二个是bool 代表key是否存在。
		fmt.Println("value=", v, "has=", has)
	} else {
		fmt.Println("false value=", v, "has=", has)
	}
}
















