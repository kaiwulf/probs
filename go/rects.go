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

import "fmt"

//     r2 = rect {
//         top_left: {-1, 3},
//         dimensions: {2, 1}
//     }

//     r3 = rect {
//         top_left: {0, 5},
//         dimensions: {4, 3}
//     }
// )

func main() {
	// r1 := rects { top_left: [1, 4], dimensions: [3, 3] }
	// m = make(map[[]int] []int)
	t := []struct {
		topleft    [2]int
		dimensions [2]int
	}{
		{[2]int{1, 4}, [2]int{2, 1}},
		{[2]int{-1, 3}, [2]int{2, 1}},
	}

	fmt.Println(t)

}

// func creat_rect(top_left [2]int, dimensions [2]int) {
//     var r = rect {
//         top_left: top_left,
//         dimensions: dimensions
//     }

//     return r
// }
