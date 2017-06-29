package main

import "fmt"



func main() {
	lable:
	for i := 1; i < 3; i++ {
		fmt.Println("outer:%d", i)
		for j := 1; j < 3; j++ {
			fmt.Println("inner:%d", j)
			if j == 2 {
				break lable
			}
		}
	}
}
