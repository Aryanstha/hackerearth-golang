package main

import (
    "fmt"
)

const mod = 1000000007

func main() {
    var t int64
    fmt.Scan(&t)
    for i := int64(0); i < t; i++ {
        var n int64
        fmt.Scan(&n)
        ans := (1 + (((2*n + 1) % mod) * ((n - 1) % mod)) % mod) % mod
        fmt.Println(ans)
    }
}
