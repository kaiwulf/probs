/*
* Euclid's algorithm for finding gcd
* On command line: euclid n m
* where n and m are integers
 */

package main

import (
	"fmt"
	"os"
	"strconv"
)

func gcd(m int64, n int64) int64 {
	r := m % n
	if r == 0 {
		return n
	} else {
		return gcd(n, r)
	}
}

func main() {
	for l, o := range os.Args {
		fmt.Println(l, o)
	}

	if len(os.Args) > 3 {
		fmt.Println("USAGE: euclid n m")
		os.Exit(3)
	}

	var err error
	N := make([]int, len(os.Args))
	for i := 1; i < len(os.Args); i++ {
		if N[i], err = strconv.Atoi(os.Args[i]); err != nil {
			panic(err)
		}

	}

	n := int64(N[1])
	m := int64(N[2])

	g := gcd(n, m)
	fmt.Println(g)
}
