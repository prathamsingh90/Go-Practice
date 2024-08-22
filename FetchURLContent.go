package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	for _, arg := range os.Args[1:] {
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Printf("Failed with error:%v\n", err)
			continue
		}
		data, err := io.ReadAll(resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed with error:%v\n", err)
			continue
		}
		fmt.Printf("Response:%s\n", data)
	}
	fmt.Printf("Execution Time:%v\n", time.Since(start).Seconds())

}
