package main

import (
    "fmt"
)

func moduloExponentiation(A, B, M int64) int64 {
    result := int64(1)
    for B > 0 {
        if B%2 == 1 {
            result = (result * A) % M
        }
        A = (A * A) % M
        B = B / 2
    }
    return result
}

func main() {
    var A, B int64
    M := int64(1000000007)
    fmt.Scan(&A, &B)
    v := moduloExponentiation(A, B, M)
    fmt.Println(v)
}
