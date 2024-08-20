package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()

	input := os.Args[1:]
	data := make(map[string]int)

	if len(input) == 0 {
		getRepeatedLine(os.Stdin, data)
	} else {
		for _, f := range input {
			file, err := os.Open(f)
			if err != nil {
				fmt.Printf("Error opening file: %v\n", err)
				continue
			}
			getRepeatedLine(file, data)
			defer file.Close()
		}
	}

	fmt.Printf("Execution Time:%v\n", time.Since(start).Seconds())

}

func getRepeatedLine(file *os.File, data map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		data[input.Text()]++
	}

	for line, count := range data {
		if count > 1 {
			fmt.Printf("'%s' repeats %d times in file %v \n", line, count, file.Name())
		}
	}
}
