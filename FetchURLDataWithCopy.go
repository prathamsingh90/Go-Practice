package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	for _, arg := range os.Args[1:] {
		if !strings.HasPrefix(arg, "http://") {
			arg = fmt.Sprint("http://", arg)
		}
		resp, err := http.Get(arg)
		if err != nil {
			fmt.Printf("Failed with error:%v\n", err)
			continue
		}
		_, err = io.Copy(os.Stdout, resp.Body)
		defer resp.Body.Close()
		if err != nil {
			fmt.Printf("Failed with error:%v\n", err)
			continue
		}
		fmt.Printf("\nStatus Code:%s\n", resp.Status)
	}
	fmt.Printf("Execution Time:%v\n", time.Since(start).Seconds())

}
