package main

import "fmt"

func solve(N int64) int64 {
    ans := int64(-1)
    factors := make([]int64, 0)
    simplified := []int64{2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 15, 18, 20, 24, 42}
    if N%2 != 0 {
        return ans
    }
    for i := 0; i < len(simplified); i++ {
        if N%simplified[i] == 0 {
            factors = append(factors, N/simplified[i])
        }
    }
    for a := 0; a < len(factors); a++ {
        for b := 0; b < len(factors); b++ {
            for c := 0; c < len(factors); c++ {
                for d := 0; d < len(factors); d++ {
                    if factors[a]+factors[b]+factors[c]+factors[d] == N {
                        if factors[a]*factors[b]*factors[c]*factors[d] > ans {
                            ans = factors[a] * factors[b] * factors[c] * factors[d]
                        }
                    }
                }
            }
        }
    }
    return ans
}

func main() {
    var testCases int
    fmt.Scan(&testCases)
    var N int64
    for testCases > 0 {
        fmt.Scan(&N)
        fmt.Println(solve(N))
        testCases--
    }
}
