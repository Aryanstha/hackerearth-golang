package main

import (
    "fmt"
    "sort"
)

func main() {
    var m, i, j int64
    fmt.Scan(&m)
    a := make([]int64, m)
    for i = 0; i < m; i++ {
        fmt.Scan(&a[i])
    }
    var f []int64
    sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
    d := a[1] - a[0]
    for i = 1; i*i <= d; i++ {
        if d%i == 0 {
            if i*i == d {
                f = append(f, i)
            } else {
                f = append(f, i)
                f = append(f, d/i)
            }
        }
    }
    sort.Slice(f, func(i, j int) bool { return f[i] < f[j] })
    for i = 0; i < int64(len(f)); i++ {
        var c int64 = 0
        for j = 2; j < m; j++ {
            if a[j]%f[i] != a[1]%f[i] {
                c = 1
                break
            }
        }
        if c == 0 && f[i] != 1 {
            fmt.Printf("%d ", f[i])
        }
    }
}
