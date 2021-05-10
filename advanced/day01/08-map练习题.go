package main

import (
	"strings"
	"fmt"
)

func wordCountFunc(str string) map[string]int {
	s := strings.Fields(str)		// 将字符串，拆分成 字符串切片s
	m := make(map[string]int)		// 创建一个用于存储 word 出现次数的 map

	// 遍历拆分后的字符串切片
	for i:=0; i<len(s);i++ {
		if _, ok := m[s[i]]; ok {			// ok == ture 说明 s[i] 这个key存在
			m[s[i]] = m[s[i]]+1			// m[s[i]]++
		} else {							// 说明 s[i] 这个key不存在， 第一次出现。添加到map中
			m[s[i]] = 1
		}
	}
	return m
}

func wordCountFunc2(str string) (m map[string]int) {
	m = make(map[string]int)
	arr := strings.Fields(str)
	for _, v := range arr {
		m[v]++
	}
	return
}

func main()  {
	str := "I love my work and I I I I love love love my family too"
	//mRet := wordCountFunc(str)

	mRet := wordCountFunc2(str)

	// 遍历map ，展示每个word 出现的次数：
	for k, v := range mRet {
		fmt.Printf("%q:%d\n", k, v)
	}
}
