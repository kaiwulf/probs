package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	for j := 1; j < len(arr); j++ {
		key := arr[j]
		i := j - 1
		fmt.Println(arr[i], key)
		for ; i > 0 && arr[i] > key; i-- {
			// swap := arr[i+1]
			arr[i+1] = arr[i]
			// arr[i] = swap
		}
		arr[i+1] = key
	}
	fmt.Println(arr)
	return arr
}

func main() {
	var a = []int{5, 2, 4, 6, 1, 3}
	var b = []int{31, 41, 59, 26, 41, 58}
	c := insertionSort(a)
	d := insertionSort(b)
	fmt.Println(c, d)
}
