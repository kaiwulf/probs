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

import "fmt"

type rects struct {
    top_left [2]int
    dimensions [2]int
}

var (
    r1 = rect {
        top_left: {1, 4},
        dimensions: {3, 3}
    }

    r2 = rect {
        top_left: {-1, 3},
        dimensions: {2, 1}
    }

    r3 = rect {
        top_left: {0, 5},
        dimensions: {4, 3}
    }
)

func main() {
    m = make(map[string] [2]int)


}

func creat_rect(top_left [2]int, dimensions [2]int) {
    var r = rect {
        top_left: top_left,
        dimensions: dimensions
    }

    return r
}

