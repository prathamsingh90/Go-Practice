package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Enter Input:")
	input := bufio.NewScanner(os.Stdin)
	arr := make([]int, 0)

	for input.Scan() {
		num, _ := strconv.ParseInt(input.Text(), 10, 64)
		arr = append(arr, int(num))
	}
	InsertionSort(&arr)
	fmt.Printf("Sorted Array is: %v\n", arr)
	fmt.Printf("Execution Time:%v\n", time.Since(start).Seconds())

}

func InsertionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		j := i
		for j > 0 && (*arr)[j-1] > (*arr)[j] {
			(*arr)[j-1], (*arr)[j] = (*arr)[j], (*arr)[j-1]
			j--
		}
	}
}
