package main

import (
	"fmt"
	"strconv"
)

func main() {
	A := []string{"1", "2", "110", "39"}
	var err error
	N := make([]int, len(A))
	for i := 0; i < len(A); i++ {
		if N[i], err = strconv.Atoi(A[i]); err != nil {
			panic(err)
		}

	}
	fmt.Println(N)
}
