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
    "sort"
)

const m int = 1000000

var arr [m+1]int
var p [m+1]int

func ispan(a int) int {
    if a%10 == 0 {
        return 0
    }
    temp1 := a
    var v []int
    for temp1 > 0 {
        k := temp1 % 10
        temp1 /= 10
        if k == 0 {
            return 0
        }
        v = append(v, k)
    }
    v = append(v, 0)
    sort.Slice(v, func(i, j int) bool {
        return v[i] < v[j]
    })
    var check bool = true
    for i := 1; i < len(v); i++ {
        if v[i]-v[i-1] != 1 {
            check = false
            break
        }
    }
    if check == false {
        return 0
    } else {
        return 1
    }
}

func panno() {
    var i int
    arr[1] = 1
    for i = 12; i <= m; i++ {
        arr[i] = ispan(i)
    }
    for i := int(1); i <= m; i++ {
        if arr[i] == 0 {
            p[i] = p[i-1]
        } else if arr[i] == 1 {
            p[i] = p[i-1] + 1
        }
    }
}

func main() {
    panno()
    var t int
    fmt.Scan(&t)
    for t > 0 {
        var l, r int
        fmt.Scan(&l, &r)
        fmt.Println(p[r] - p[l-1])
        t--
    }
}
