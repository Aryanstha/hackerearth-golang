/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var t int
    fmt.Scan(&t)

    in := bufio.NewReader(os.Stdin)
    out := bufio.NewWriter(os.Stdout)
    defer out.Flush()

    for i := 0; i < t; i++ {
        var n int
        fmt.Fscan(in, &n)

        a := make([]int64, n)
        for j := 0; j < n; j++ {
            fmt.Fscan(in, &a[j])
        }

        if n == 1 {
            fmt.Fprintln(out, a[0])
        } else {
            lc := (a[0] * a[1]) / gcd(a[0], a[1])
            for j := 2; j < n; j++ {
                lc = (lc * a[j]) / gcd(lc, a[j])
            }
            fmt.Fprintln(out, lc)
        }
    }
}

func gcd(a, b int64) int64 {
    if b == 0 {
        return a
    }
    return gcd(b, a%b)
}

