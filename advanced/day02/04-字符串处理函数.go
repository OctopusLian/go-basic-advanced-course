package main

import (
	"strings"
	"fmt"
)

func main()  {

	str := "I love my work and I love my family too"

	// 字符串按 指定分割符拆分
	ret := strings.Split(str, " I")
	for _, s := range ret {
		fmt.Println(s)
	}

	// 字符串按 空格拆分
	ret = strings.Fields(str)
	for _, s := range ret {
		fmt.Println(s)
	}

	// 判断字符串结束标记
	flg := strings.HasSuffix("test.abc", ".mp3")
	fmt.Println(flg)

	// 判断字符串起始标记
	flg = strings.HasPrefix("test.abc", "tes.")
	fmt.Println(flg)

}
