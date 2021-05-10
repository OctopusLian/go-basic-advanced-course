package main

import (
	"os"
	"fmt"
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

	n, err := f.WriteString("helloworld\r\n")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}
	fmt.Println("WriteString n = ", n)

	off, _ := f.Seek(-5, io.SeekEnd)
	fmt.Println("off:", off)

	n, _ = f.WriteAt([]byte("1111"), off)
	fmt.Println("WriteAt n :", n)
}
