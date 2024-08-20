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
	fmt.Println("Enter elements for bubble sort")
	arr := make([]int, 0)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		num, _ := strconv.ParseInt(input.Text(), 10, 64)
		arr = append(arr, int(num))
	}

	bubbleSort(&arr)
	fmt.Printf("Sorted Array after bubble sort is: %v\n", arr)
	fmt.Printf("Time Taken:%v\n", time.Since(start).Seconds())
}

func bubbleSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}
