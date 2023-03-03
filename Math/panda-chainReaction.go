package main

import (
    "fmt"
)

const mod = 1000003

func main() {
    var num int
    arr := make([]int, 1000004)
    arr[0] = 1
    for i := 1; i < 1000004; i++ {
        arr[i] = (arr[i-1] * i) % mod
    }
    fmt.Scan(&num)
    for num > 0 {
        var n, x uint64
        fmt.Scan(&n, &x)
        if n >= 1000003 {
            n = 0
        } else {
            n = x * uint64(arr[n])
        }
        fmt.Println(n % mod)
        num--
    }
}
