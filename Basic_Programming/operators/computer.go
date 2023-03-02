/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import (
    "fmt"
    "math"
)

func main() {
    var t uint64
    fmt.Scanf("%d", &t)
    for t > 0 {
        var n, ans, i uint64
        fmt.Scanf("%d", &n)
        i = 1
        for i <= uint64(math.Sqrt(float64(n))) {
            i *= 2
        }
        if n/i >= i/2 {
            ans = n - n/i
        } else {
            ans = (n - i/2) + 1
        }
        fmt.Println(ans)
        t--
    }
}
