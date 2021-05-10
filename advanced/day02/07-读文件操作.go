package main

import (
	"os"
	"fmt"
	"bufio"
	"io"
)

func main()  {
	f, err := os.OpenFile("C:/itcast/testFile.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()
	fmt.Println("successful")

	// 创建一个带有缓冲区(用户缓冲)的 reader
	reader := bufio.NewReader(f)
	for {
		buf, err := reader.ReadBytes('\n')		// 读一行数据
		if err != nil && err == io.EOF {
			fmt.Println("文件读取完毕")
			return
		} else if err != nil {
			fmt.Println("ReadBytes err:", err)
		}
		fmt.Print(string(buf))
	}
}
