package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	result := JoinStrings()
	fmt.Println(result)

	result = JoinStringsWithBuffer()
	fmt.Println(result)

}

// JoinStrings
func JoinStrings() string {
	start := time.Now()
	result := strings.Join(os.Args[1:], " ")
	fmt.Printf("Execution time for JoinStrings: %v\n", time.Since(start).Seconds())
	return result
}

// JoinStringsWithBuffer
func JoinStringsWithBuffer() string {
	start := time.Now()
	var buffer strings.Builder
	for _, arg := range os.Args[1:] {
		buffer.WriteString(arg + " ")
	}
	result := buffer.String()
	fmt.Printf("Execution time for JoinStringsWithBuffer: %v\n", time.Since(start).Seconds())
	return result
}
