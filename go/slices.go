package main

import (
	"fmt"
)

func main() {
	const idx = 12
	var initial = [idx]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 2, 2}
	var final [idx]int
	var lim = cap(initial) - 1

	for i := 0; i < lim; i++ {
		// final[lim-2 : lim-1]
		fmt.Println(initial[i : i+1])
	}
	fmt.Println(final)
}
