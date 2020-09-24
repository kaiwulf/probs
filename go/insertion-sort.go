package main

import (
    "fmt"
 )

func insertionSort(arr []int) []int {
    // const capacity = cap(arr)
    // var sorted [capacity]int
    // i := 0
    for j := 2; j < len(arr); j++ {
        key := arr[j]
        // insert arr[j] into the sorted sequence sorted[1..j-1]
        // sorted[i] = arr[j]
        i = j-1
        for ; i > 0 && arr[i] > key; i-- {
            arr[i+1] = arr[i]
        }
        arr[i+1] = key
    }
    return arr
}

func main() {
    var a = []int{5,2,4,6,1,3}
    var b = []int{31,41,59,26,41,58}
    c := insertionSort(a)
    d := insertionSort(b)
    fmt.Println(c, d)
}