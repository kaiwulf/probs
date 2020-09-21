package main
import (
    "fmt"
)

type stack struct {
    data int
    next *int
    // prev *int ??
}

func main() {
    const idx = 12
    var stack [idx]int
    var initial = [idx]int{1,1,1,1,0,0,0,0,1,1,2,2}
    var final [idx]int
    var limit1 = cap(initial) - 1 // 31
    var limit2 = cap(final) - 1   // 31

    for i := limit1-1; i >= 0; i-- {
        stack[i] = initial[limit1 - i]
        fmt.Println("stack[i] = ", stack[i])
        fmt.Println("initial[limit1-i] = ", initial[limit1 - i])
    }
    for j := ; j <= limit2; j++ {
        final[limit2 - j] = stack[j]
    }

    // fmt.Println("initial = ", initial)
    // fmt.Println("final = ", final)

}