package main

import (
    "bufio"
    "fmt"
    "os"
)

var m int = 1000000007

func fastmod(a, b, m int) int {
    ans := int(1)
    a = a % m
    if a == 0 {
        return 0
    }
    for b > 0 {
        if b&1 != 0 {
            ans = (ans * a) % m
        }
        b = b >> 1
        a = (a * a) % m
    }
    return ans
}

func inv(a, m int) int {
    return fastmod(a, m-2, m)
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Split(bufio.ScanWords)
    scanner.Scan()
    t := toInt(scanner.Text())

    for i := 0; i < t; i++ {
        scanner.Scan()
        n := toInt(scanner.Text())
        ans1 := (((n % m) * (n - 1) % m) * (n - 1) % m * inv(4, m) % m) % m
        ans2 := ((n % m) * (n - 1) % m * (2*n - 1) % m * inv(6, m)) % m
        fmt.Printf("%d %d\n", ans1, ans2)
    }
}

func toInt(s string) int {
    var res int
    for i := 0; i < len(s); i++ {
        res = res*10 + int(s[i]-'0')
    }
    return res
}
