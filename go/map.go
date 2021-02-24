package main 

import (
    "fmt"
)

func main() {
    var t = make(map[[]int] []int)

    t[[]int{91,62}] = []int{4,5}

    fmt.Println(t)
}