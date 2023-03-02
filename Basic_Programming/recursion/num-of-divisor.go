package main

import "fmt"

const mod = 1000000007

func recur(sum, N, k int64) int64 {
    sum =  sum - k*(N*(N + 1)/2) + (N*(N + 1)/ 2)
    N = N / k
    if N <= 0 {
        return sum
    }
    return recur(sum, N, k)
}

func main() {
    var t int
  var N int64
    fmt.Scan(&t)
    for i := 0; i < t; i++ {
        var n, k int64
        fmt.Scan(&n, &k)
        sum := (n*(n+1))/2
        N = n / k
        fmt.Println(recur(sum, N, k))
    }
}
