package main

import (
    "fmt"
)

func main() {
    const idx = 7
    var one = [idx]int{1,2,3,4,5,6,7}
    var two = [idx]int{10, 9, 8, 7, 6, 5, 4}
    var limit2 = cap(two)
    
    for i:= 0; i < cap(one); i++ {
        fmt.Println("one[i] = ", one[i])
    }
    for j := limit2 - 1; j >= 0; j-- {
        fmt.Println("j = ", j)
        fmt.Println("two[j] = ", two[j])
    }
}
