package main

import (
	"strings"
	"fmt"
	"os"
	"bufio"
	"io"
)

// 从一个文件中逐行读取内容，统计该文件共有多少个 love
func readFile(fileName string) int  {
	fp, err := os.Open(fileName)
	if err !=nil{
		fmt.Println("Open err：", err)
		return -1
	}
	defer fp.Close()

	row := bufio.NewReader(fp)	// 创建一个reader
	var total int = 0			// 统计 Love 总数的变量

	for {	// 循环 按行读取文件内容，存入buf中
		buf, err :=row.ReadBytes('\n')
		if err != nil && err == io.EOF {
			break
		}
		// 交给 wordCount 统计每行中 love 出现的次数。累加各行 love 数
		total += wordCount(string(buf[:]))
	}
	return total
}

// 从一行字符串中获取 love 出现的次数。将该次数返回。
func wordCount(s string) int {
	arr := strings.Fields(s)			// 分割字符串,存入字符数组
	m := make(map[string]int)			// 创建map

	// 对arr中的每个单词进行循环，存入map中，统计
	for i:= 0; i<len(arr); i++ {
		m[arr[i]]++
	}
	// 统计 map 中一共有多少个 "Love"
	for k, v := range m {
		if k == "love" {
			fmt.Printf("%s : %d\n", k, v)
			return m[k]			// 返回 Love 出现的次数
		}
	}
	return 0		// 没有Love， 返回 0
}

func main()  {
/*	// 测试从文件读一行
	oneFilenum := readFile("C:/itcast/test2/test.txt")
	fmt.Printf("一个文件中有：%d 个Love\n", oneFilenum)

	writeString
*/

	fmt.Println("请输入要找寻的目录：")
	var path string
	fmt.Scan(&path)						// 获取用户指定的目录名

	dir, _ := os.OpenFile(path, os.O_RDONLY, os.ModeDir)	// 只读打开该目录
	defer dir.Close()

	names, _ := dir.Readdir(-1)		// 读取当前目录下所有的文件名和目录名，存入names切片

	var AllLove int = 0
	for _, name := range names {		// 遍历切片，获取文件/目录名
		if !name.IsDir() {
			s := name.Name()			// 文件名不包含路径。
			if strings.HasSuffix(s,".txt") {
				AllLove += readFile(path + s)		// 拼接有路径的文件名（绝对路径）
			}
		}
	}
	fmt.Printf("目录所有文件中共有 %d 个Love\n", AllLove)
}
