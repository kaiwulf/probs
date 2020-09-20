/*
*    You are given given a list of rectangles represented by min and max x- and y-coordinates. Compute whether or not a pair of rectangles overlap each other. If one rectangle completely covers another, it is considered overlapping.
*
*    For example, given the following rectangles:
*
*    {
*        "top_left": (1, 4),
*        "dimensions": (3, 3) # width, height
*    },
*    {
*        "top_left": (-1, 3),
*        "dimensions": (2, 1)
*    },
*    {
*        "top_left": (0, 5),
*        "dimensions": (4, 3)
*    }
*    return true as the first and third rectangle overlap each other.
 */

package main

import (
	"fmt"
	"math"
)

type rect struct {
	topleft    [2]int
	dimensions [2]int
}

func main() {
	// m = make(map[[2]int][]int)

	// t := []struct {
	// 	topleft    [2]int
	// 	dimensions [2]int
	// }{
	t := []rect{
		{[2]int{1, 4}, [2]int{2, 1}},
		{[2]int{-1, 3}, [2]int{2, 1}},
		{[2]int{0, 5}, [2]int{4, 3}},
	}

	fmt.Println(t)
	fmt.Println(dist(t[0]))

}

func dist(d rect) float64 {
	var x1, x2 float64 = float64(d.topleft[0]), float64(d.topleft[1])
	fmt.Printf("%f %f", x1, x2)
	return math.Abs(x1 - x2)
}

// func creat_rect(top_left [2]int, dimensions [2]int) {
//     var r = rect {
//         top_left: top_left,
//         dimensions: dimensions
//     }

//     return r
// }
