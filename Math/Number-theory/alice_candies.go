package main

import (
    "fmt"
    )
func main(){
    
var n int64
fmt.Scanln(&n)

cnt := 0

for i := int64(1); i*i < (n + 5); i++ {

    if (n-i*(i-1))%i == 0 {

        res := (n - i*(i-1)) / i

        if res%2 == 1 {

            cnt++
        }
    }
}

fmt.Println(cnt)
}