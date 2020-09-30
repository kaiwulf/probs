package main

import (
	"fmt"
)

func main() {
	var inp int
	fmt.Println("Enter array limit: ")
	fmt.Scanf("%d", inp)

	const idx = inp
	// var initial = [idx]int{1, 1, 1, 1, 0, 0, 0, 0, 1, 1, 2, 2}
	var initial = [idx]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var final [idx]int
	var limit = cap(initial) - 1

	for i := limit; i >= 0; i-- {
		final[i] = initial[limit-i]
		// fmt.Printf("final[%d] = %d\ninitial[%d] = %d\n", i, final[i], limit-i, initial[limit-i])
	}

	fmt.Println("initial = ", initial)
	fmt.Println("final = ", final)

}
