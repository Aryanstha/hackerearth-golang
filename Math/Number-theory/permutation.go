package main

import (
    "fmt"
)

var ar [100001]int64

func permutationSum() {
    ar[0] = 0
    ar[1] = 1
    ar[2] = 1
    for i := int64(3); i < 100001; i++ {
        ar[i] = i - (i % 2) + ar[i-1]
    }
}

func main() {
    permutationSum()
    var t int64
    fmt.Scan(&t)
    for i := int64(0); i < t; i++ {
        var n int64
        fmt.Scan(&n)
        fmt.Println(ar[n])
    }
}
