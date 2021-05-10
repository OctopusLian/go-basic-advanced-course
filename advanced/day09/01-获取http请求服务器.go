package main

import (
	"net"
	"fmt"
	"os"
)

func errFunc(err error, info string)  {
	if err != nil {
		fmt.Println(info, err)
		//return					// 返回当前函数调用
		//runtime.Goexit()			// 结束当前go程
		os.Exit(1)			// 将当前进程结束。
	}
}

func main()  {

	listener, err := net.Listen("tcp", "127.0.0.1:8000")
	errFunc(err, "net.Listen err:")
	defer listener.Close()

	conn, err := listener.Accept()
	errFunc(err, "Accpet err:")
	defer conn.Close()

	buf := make([]byte, 4096)
	n, err := conn.Read(buf)
	if n == 0 {
		return
	}
	errFunc(err, "conn.Read")

	fmt.Printf("|%s|\n", string(buf[:n]))
}
