/*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here
package main

import("fmt")

func main(){
      var n int
  
  var fact int = 1


  
  fmt.Scan(&n)


  for i := 1; i <= n; i++ {
      
    fact = fact * i
    
  }

  fmt.Println(fact)
  
}
    

