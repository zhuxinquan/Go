package main

import (
	"fmt"
	"os"
)

func main() {
	userFile := "astaxie.txt"
	fout, err := os.Create(userFile)
	if err != nil {
		fmt.Println(userFile, err)
		return
	}
	defer fout.Close()
	for i := 0; i < 10; i++ {
		fout.WriteString("Just a test!\r\n")
		fout.Write([]byte("Just a test!\r\n"))
	}

	//在该位置以覆盖的方式写入
	fout.WriteAt([]byte("123"), 4)
}