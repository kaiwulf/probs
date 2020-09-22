package main

import (
	"fmt"
)

func main() {
	const idx = 12
	// var stack [idx]int
	var initial = [idx]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 2, 2}
	var final [idx]int
	var limit1 = cap(initial) - 1 // 11
	// var limit2 = cap(final) - 1   // 11

	for i := limit1; i >= 0; i-- {
		final[i] = initial[limit1-i]
		fmt.Printf("final[%d] = %d\ninitial[%d] = %d\n", i, final[i], limit1-i, initial[limit1-i])
	}

	fmt.Println("initial = ", initial)
	fmt.Println("final = ", final)

}
