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
	fmt.Println("Enter elements for selection sort")
	arr := make([]int, 0)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		num, _ := strconv.ParseInt(input.Text(), 10, 64)
		arr = append(arr, int(num))
	}

	selectionSort(&arr)
	fmt.Printf("Sorted Array after selection sort is: %v\n", arr)
	fmt.Printf("Time Taken:%v\n", time.Since(start).Seconds())
}

func selectionSort(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minIndex := i
		for j := i; j < len(*arr); j++ {
			if (*arr)[minIndex] > (*arr)[j] {
				minIndex = j
			}
		}
		(*arr)[i], (*arr)[minIndex] = (*arr)[minIndex], (*arr)[i]
	}
}
