package main

import (
	"os"
	"fmt"
	"io"
)

func main()  {

	// 打开读文件
	f_r, err := os.Open("C:/itcast/01-复习.avi")
	if err != nil {
		fmt.Println("Open err: ", err)
		return
	}
	defer f_r.Close()
	// 创建写文件
	f_w, err := os.Create("C:/itcast/test.avi")
	if err != nil {
		fmt.Println("Create err: ", err)
		return
	}
defer f_w.Close()

	// 从读文件中获取数据，放到缓冲区中。
	buf := make([]byte, 4096)
	// 循环从读文件中，获取数据，“原封不动的”写到写文件中。
	for {
		n, err := f_r.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Printf("读完。n = %d\n", n)
			return
		}
		f_w.Write(buf[:n])		// 读多少，写多少
	}

}
