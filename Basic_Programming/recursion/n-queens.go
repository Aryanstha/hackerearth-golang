package main

import (
    "fmt"
)

func issafe(arr [][]int, x, y, n int) bool {
    for row := 0; row < x; row++ {
        if arr[row][y] == 1 {
            return false
        }
    }

    row := x
    col := y

    for row >= 0 && col >= 0 {
        if arr[row][col] == 1 {
            return false
        }

        row--
        col--
    }

    row = x
    col = y

    for row >= 0 && col < n {
        if arr[row][col] == 1 {
            return false
        }

        row--
        col++
    }

    return true
}

func nqueen(arr [][]int, x, n int) bool {
    if x >= n {
        return true
    }

    for col := 0; col < n; col++ {
        if issafe(arr, x, col, n) {
            arr[x][col] = 1

            if nqueen(arr, x+1, n) {
                return true
            }

            arr[x][col] = 0
        }
    }

    return false
}

func main() {
    var n int
    fmt.Scan(&n)

    arr := make([][]int, n)

    for i := 0; i < n; i++ {
        arr[i] = make([]int, n)

        for j := 0; j < n; j++ {
            arr[i][j] = 0
        }
    }

    if nqueen(arr, 0, n) {
        for i := 0; i < n; i++ {
            for j := 0; j < n; j++ {
                fmt.Printf("%d ", arr[i][j])
            }
            fmt.Println()
        }
    } else {
        fmt.Println("Not possible")
    }
}
