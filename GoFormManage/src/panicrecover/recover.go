package main

import (
	"fmt"
	"os"
)

func main() {
	defer func(){
		fmt.Println("recover")
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	fmt.Println("panic")
	//panic("err")
	file, err := os.OpenFile("1234", os.O_APPEND|os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	fmt.Println("before close")
	file.Close()
	fmt.Println("after")
}
