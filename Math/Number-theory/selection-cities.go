package main

import "fmt"

const mod = 1000000007

func frac(a, m int64) int64 {
    s := int64(1)
    for m > 0 {
        if m%2 == 1 {
            s = (s * a) % mod
        }
        a = (a * a) % mod
        m /= 2
    }
    return s % mod
}

func main() {
    var q int64
    fmt.Scan(&q)
    for q > 0 {
        var m int64
        fmt.Scan(&m)
        ans := frac(2, m) - 1
        if ans < 0 {
            ans += mod
        }
        fmt.Println(ans)
        q--
    }
}
