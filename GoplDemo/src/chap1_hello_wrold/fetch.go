package main

import (
	"time"
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	for i := 0; i < 5; i++ {
		start := time.Now()
		resp, err := http.Get(os.Args[1])
		if err != nil {
			fmt.Println(err) // send to channel ch
			os.Exit(-1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		fmt.Println(string(b))
		fmt.Printf("%.2fs\tbytes: %d\n", time.Since(start).Seconds(), len(b))
	}
}