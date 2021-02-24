package main

import (
	"fmt"
)

func insertionSort(arr []int) []int {
	fmt.Println(arr)
	var i int
	for j := 0; j < len(arr); j++ {
		key := arr[j]
		if j > 0 {
			i = j - 1
		}
		// fmt.Println(i, key)
		// var step rune
		// fmt.Println(arr[i], key)
		for k := 1; i > 0 && arr[i] > key; i-- {
			fmt.Println("i", i)
			fmt.Println("k", k)
			// fmt.Printf("before assignment:\narr[%d] = %d\n", i+1, arr[i])
			// fmt.Scanf("%r", step)
			arr[i+1] = arr[i]
			k++
			// fmt.Printf("after assignment:\narr[%d] = %d\n", arr[i+1], arr[i])
			// fmt.Scanf("%r", step)
		}
		arr[i+1] = key
	}
	// fmt.Println(arr)
	return arr
}

func main() {
	var a = []int{5, 2, 4, 6, 1, 3}
	var b = []int{31, 41, 59, 26, 41, 58}
	c := insertionSort(a)
	d := insertionSort(b)
	fmt.Println(c, d)
}
