package main

import (
    "fmt"
)

func gcd(a, b int64) int64 {
    if a == 0 {
        return b
    }
    return gcd(b%a, a)
}

func findGCD(arr []int64, n int64) int64 {
    result := arr[0]
    for i := int64(1); i < n; i++ {
        result = gcd(arr[i], result)
        if result == 1 {
            return 1
        }
    }
    return result
}

func exponentiation(prod, y int64) int64 {
    resultt := int64(1)
    M := int64(1000000007)
    for y > 0 {
        if y%2 == 1 {
            resultt = (resultt * prod) % M
        }
        prod = (prod * prod) % M
        y = y / 2
    }
    return resultt
}

func main() {
    var n int64
    var M int64 = 1000000007
    fmt.Scan(&n)
    arr := make([]int64, n)
    var prod int64 = 1
    var ans int64
    for i := int64(0); i < n; i++ {
        fmt.Scan(&arr[i])
    }
    for i := int64(0); i < n; i++ {
        prod = (prod % M * arr[i] % M) % M
    }
    y := findGCD(arr, n)
    ans = exponentiation(prod, y)
    fmt.Println(ans)
}
