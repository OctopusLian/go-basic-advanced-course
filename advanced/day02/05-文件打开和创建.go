package main

import (
	"os"
	"fmt"
)

/*func main()  {
	f, err := os.Create("C:/itcast/testFile.xyz")
	if err != nil {
		fmt.Println("create err:", err)
		return
	}
	defer f.Close()
	fmt.Println("successful")

}*/

/*func main()  {
	f, err := os.Open("C:/itcast/testFile.xyz")
	if err != nil {
		fmt.Println("open err:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("######")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}

	fmt.Println("successful")
}*/


func main()  {

	f, err := os.OpenFile("C:/itcast/testFile.xyz", os.O_RDWR, 6)
	if err != nil {
		fmt.Println("OpenFile err:", err)
		return
	}
	defer f.Close()

	_, err = f.WriteString("######")
	if err != nil {
		fmt.Println("WriteString err:", err)
		return
	}

	fmt.Println("successful")
}
